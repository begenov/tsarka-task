package service

import (
	"context"
	"database/sql"

	"github.com/begenov/tsarka-task/internal/domain"
	"github.com/begenov/tsarka-task/internal/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user domain.User) (int, error) {
	id, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, domain.NotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, domain.NotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NotFound
		}
		return err
	}

	return nil
}
