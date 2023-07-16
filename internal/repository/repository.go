package repository

import (
	"context"

	"github.com/begenov/tsarka-task/internal/domain"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go
type Users interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Counters interface {
	Add(key string, value int64) (int64, error)
	Sub(key string, value int64) (int64, error)
	Get(key string) (int64, error)
}
