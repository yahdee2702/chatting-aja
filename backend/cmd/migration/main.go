package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yahdee2702/chatting-aja/internal/config"
	"github.com/yahdee2702/chatting-aja/internal/database"
	"github.com/yahdee2702/chatting-aja/internal/logging"
)

const migrationsTable = "migrations"

type migrationFile struct {
	version   int
	name      string
	direction string
	path      string
	applied   bool
	batch     int
}

type migrationFileInfo struct {
	applied bool
	batch   int
}

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	logger := logging.New(config.App.Env)
	slog.SetDefault(logger)

	ctx := context.Background()

	db, err := database.NewPostgres(config.DB)

	if err != nil {
		logger.ErrorContext(ctx, "failed to connect to database", "error", err)
		return
	}

	if err = ensureMigrationTable(ctx, db); err != nil {
		logger.ErrorContext(ctx, "failed to create migration table", "error", err)
		return
	}

	migrationsDir := getMigrationsPath()

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		logger.ErrorContext(ctx, "failed to read migrations dir", "error", err)
		return
	}

	databaseInfo, err := fetchDatabaseInfo(ctx, db)
	if err != nil {
		logger.ErrorContext(ctx, "failed to fetch database migration file info", "error", err)
		return
	}

	upFiles, downFiles := fetchMigrationFiles(migrationsDir, files, databaseInfo)

	var cmd string
	if len(os.Args) < 2 {
		cmd = "up"
	} else {
		cmd = os.Args[1]
	}

	switch cmd {
	case "up":
		if err = runUp(ctx, logger, db, upFiles); err != nil {
			logger.ErrorContext(ctx, "failed to migrate up", "error", err)
			return
		}
	case "down":
		if err = runDown(ctx, logger, db, downFiles); err != nil {
			logger.ErrorContext(ctx, "failed to migrate up", "error", err)
			return
		}
	}
}

func ensureMigrationTable(ctx context.Context, db *sqlx.DB) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version BIGINT PRIMARY KEY,
			batch INT NOT NULL,
			name TEXT NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`, migrationsTable))

	return err
}

func getMigrationsPath() string {
	candidates := []string{
		"migrations",
		"../migrations",
		"../../migrations",
	}

	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			return c
		}
	}

	exe, _ := os.Executable()
	return filepath.Join(filepath.Dir(exe), "migrations")
}

func fetchMigrationFiles(migrationsDir string, files []os.DirEntry, databaseInfo map[int]migrationFileInfo) (upFiles []migrationFile, downFiles []migrationFile) {
	for _, file := range files {
		mFile, ok := parseMigrationFile(migrationsDir, file.Name())
		if !ok {
			continue
		}

		if fileInfo, ok := databaseInfo[mFile.version]; ok {
			mFile.applied = fileInfo.applied
			mFile.batch = fileInfo.batch
		}

		if mFile.direction == "up" {
			upFiles = append(upFiles, mFile)
		} else {
			downFiles = append(downFiles, mFile)
		}
	}

	sort.Slice(upFiles, func(i, j int) bool { return upFiles[i].version < upFiles[j].version })
	sort.Slice(downFiles, func(i, j int) bool { return downFiles[i].version > downFiles[j].version })

	return upFiles, downFiles
}

func fetchDatabaseInfo(ctx context.Context, db *sqlx.DB) (map[int]migrationFileInfo, error) {
	rows, err := db.QueryxContext(ctx, fmt.Sprintf("SELECT version, batch FROM %s ORDER BY version ASC", migrationsTable))
	if err != nil {
		return nil, fmt.Errorf("sql query: %w", err)
	}
	defer rows.Close()

	migrationFileInfos := make(map[int]migrationFileInfo)

	for rows.Next() {
		var version int
		var batch int
		if err := rows.Scan(&version, &batch); err != nil {
			return nil, fmt.Errorf("row scan: %w", err)
		}

		migrationFileInfos[version] = migrationFileInfo{
			applied: true,
			batch:   batch,
		}
	}

	return migrationFileInfos, nil
}

func parseMigrationFile(migrationsDir string, filename string) (migrationFile, bool) {
	parts := strings.Split(filename, ".")
	if len(parts) < 3 || parts[len(parts)-1] != "sql" {
		return migrationFile{}, false
	}

	direction := parts[len(parts)-2]
	if direction != "up" && direction != "down" {
		return migrationFile{}, false
	}

	base := strings.Join(parts[:len(parts)-2], ".")
	underIdx := strings.Index(base, "_")
	if underIdx < 0 {
		return migrationFile{}, false
	}

	versionStr := base[:underIdx]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return migrationFile{}, false
	}

	return migrationFile{
		version:   version,
		name:      base[underIdx+1:],
		direction: direction,
		path:      filepath.Join(migrationsDir, filename),
		batch:     0,
		applied:   false,
	}, true
}

func runUp(ctx context.Context, logger *slog.Logger, db *sqlx.DB, files []migrationFile) error {
	currentBatch := 0
	for _, file := range files {
		if file.applied {
			currentBatch = file.batch
			continue
		}

		file.batch = currentBatch + 1

		logger.Info(fmt.Sprintf("applying %06d_%s", file.version, file.name))

		if err := execMigration(ctx, db, file); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("failed applying %06d_%s", file.version, file.name))
			return fmt.Errorf("migration %06d_%s: %w", file.version, file.name, err)
		}

		if err := logMigration(ctx, db, file); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("failed logging %06d_%s", file.version, file.name))
			return fmt.Errorf("logging migration %06d_%s: %w", file.version, file.name, err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("successfully migrating %06d_%s", file.version, file.name))
	}

	return nil
}

func runDown(ctx context.Context, logger *slog.Logger, db *sqlx.DB, files []migrationFile) error {
	currentBatch := 0
	for _, file := range files {
		if !file.applied {
			continue
		}

		if currentBatch == 0 {
			currentBatch = file.batch
		}

		if file.batch != currentBatch {
			return nil
		}

		logger.Info(fmt.Sprintf("rolling back %06d_%s", file.version, file.name))

		if err := execMigration(ctx, db, file); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("failed rolling back %06d_%s", file.version, file.name))
			return fmt.Errorf("rollback %06d_%s: %w", file.version, file.name, err)
		}

		if err := dropMigration(ctx, db, file); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("failed droppping %06d_%s", file.version, file.name))
			return fmt.Errorf("dropping rollback %06d_%s: %w", file.version, file.name, err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("successfully rolling back %06d_%s", file.version, file.name))
	}

	return nil
}

func execMigration(ctx context.Context, db *sqlx.DB, f migrationFile) error {
	sql, err := os.ReadFile(f.path)

	if err != nil {
		return fmt.Errorf("file reading: %w", err)
	}

	if _, err = db.ExecContext(ctx, string(sql)); err != nil {
		return err
	}

	return err
}

func logMigration(ctx context.Context, db *sqlx.DB, f migrationFile) error {
	if _, err := db.ExecContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s (version, batch, name) VALUES ($1,$2,$3)", migrationsTable),
		f.version, f.batch, f.name,
	); err != nil {
		return err
	}

	return nil
}

func dropMigration(ctx context.Context, db *sqlx.DB, f migrationFile) error {
	if _, err := db.ExecContext(
		ctx,
		fmt.Sprintf("DELETE FROM %s WHERE version = $1", migrationsTable),
		f.version,
	); err != nil {
		return err
	}

	return nil
}
