package repository

import (
	"context"

	"github.com/sumukhj1219/go-profiler/internals/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, user models.User) error
	GetUserById(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context, user models.User) error
}
