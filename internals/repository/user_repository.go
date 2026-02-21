package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sumukhj1219/go-profiler/internals/models"
)

var ErrDuplicateEmailOrUsername = errors.New("email or username already exists")
var ErrNotFound = errors.New("no user found")

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user models.User) error {
	query := createUserQuery
	_, err := u.db.NamedExecContext(ctx, query, user)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateEmailOrUsername
		}
		return err
	}
	return nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, getUserByEmailQuery, email)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, getUserByIDQuery, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &user, nil
}

func (u *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	err := u.db.SelectContext(ctx, &users, getUsersQuery)
	if err != nil {
		return nil, err
	}
	if users == nil {
		users = make([]models.User, 0)
	}

	return users, nil
}
