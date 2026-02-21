package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sumukhj1219/go-profiler/internals/models"
	"github.com/sumukhj1219/go-profiler/internals/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
	redisClient    *redis.Client
}

func NewService(userRepository repository.UserRepository, redisClient *redis.Client) UserService {
	return &userService{
		userRepository: userRepository,
		redisClient:    redisClient,
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

	cacheKey := fmt.Sprintf("user:email:%s", email)

	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(val), &user)
		return &user, nil
	}

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return user, err
}

func (s *userService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	cacheKey := fmt.Sprintf("user:id:%d", id)

	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(val), &user)
		return &user, nil
	}

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
	cacheKey := "users:all"

	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var users []models.User
		json.Unmarshal([]byte(val), &users)
		return &users, nil
	}

	if err != nil && err != redis.Nil {
		return nil, err
	}

	users, err := s.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersJSON, _ := json.Marshal(users)
	s.redisClient.Set(ctx, cacheKey, usersJSON, 2*time.Minute)

	return &users, nil
}
