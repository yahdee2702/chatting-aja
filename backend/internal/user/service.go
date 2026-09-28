package user

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetUsers(ctx context.Context, limit int) ([]User, error) {
	if limit <= 0 {
		limit = 1000
	}

	return s.repository.GetUsersWithLimit(ctx, limit)
}
