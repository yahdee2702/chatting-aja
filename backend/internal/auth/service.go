package auth

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
	jwt  *JwtHandler
}

func NewService(repo *Repository, jwt *JwtHandler) *Service {
	return &Service{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)

	if err != nil {
		return "", fmt.Errorf("user with this email is not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))

	if err != nil {
		return "", fmt.Errorf("password is incorrect")
	}

	token, err := s.jwt.Generate(user.Id)
	if err != nil {
		return "", fmt.Errorf("cannot generate jwt token: %w", err)
	}

	return token, nil
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (string, error) {
	if req.Password != req.PasswordConfirm {
		return "", errors.New("confirmed password needs to match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("error when generating password: %w", err)
	}

	userId, err := s.repo.CreateUser(ctx, req.Name, req.Email, string(passwordHash))
	if err != nil {
		return "", fmt.Errorf("error when creating user: %w", err)
	}

	token, err := s.jwt.Generate(userId)
	if err != nil {
		return "", fmt.Errorf("cannot generate jwt token: %w", err)
	}

	return token, nil
}
