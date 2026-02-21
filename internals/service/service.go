package service

import (
	"context"
	"errors"

	"github.com/sumukhj1219/go-profiler/internals/models"
)

var (
	ErrDuplicateEmailOrUsername = errors.New("email or username already exists")
	ErrInvalidEmailFormat       = errors.New("email address is invalid")
	ErrNotFound                 = errors.New("user not found")
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	GetUsers(ctx context.Context) (*[]models.User, error)
}
