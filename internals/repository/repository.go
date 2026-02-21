package repository

import (
	"context"

	"github.com/sumukhj1219/go-profiler/internals/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
