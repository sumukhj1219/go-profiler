package service

import (
	"context"
	"errors"
	"net/mail"

	"github.com/sumukhj1219/go-profiler/internals/models"
	"github.com/sumukhj1219/go-profiler/internals/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	err = s.userRepository.CreateUser(ctx, user)
	if errors.Is(err, repository.ErrDuplicateEmailOrUsername) {
		return ErrDuplicateEmailOrUsername
	}
	return err
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmailFormat
	}
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return user, err
}

func (s *userService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.userRepository.GetUserById(ctx, id)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(ctx context.Context) (*[]models.User, error) {
	users, _ := s.userRepository.GetUsers(ctx)
	return &users, nil
}
