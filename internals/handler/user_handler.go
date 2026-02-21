package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/sumukhj1219/go-profiler/internals/models"
	"github.com/sumukhj1219/go-profiler/internals/service"
	"github.com/sumukhj1219/go-profiler/pkg/logger"
)

type userHandler struct {
	userService service.UserService
	logger      logger.Logger
}

func NewHandler(userService service.UserService, logger logger.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

type UpdateUserRequest struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Phone        string     `json:"phone"`
	ProfileImage string     `json:"profile_image"`
	Bio          string     `json:"bio"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	IsActive     bool       `json:"is_active"`
	IsVerified   bool       `json:"is_verified"`
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.userService.CreateUser(r.Context(), user); err != nil {
		if errors.Is(err, service.ErrDuplicateEmailOrUsername) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		h.logger.Error("failed to create user", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := models.User{
		ID:           req.ID,
		Username:     req.Username,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		ProfileImage: req.ProfileImage,
		Bio:          req.Bio,
		DateOfBirth:  req.DateOfBirth,
		IsActive:     req.IsActive,
		IsVerified:   req.IsVerified,
	}

	if err := h.userService.UpdateUser(r.Context(), user); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByEmail(r.Context(), email)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		h.logger.Error("failed to get user by email", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserById(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		h.logger.Error("failed to get user by id", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers(r.Context())
	if err != nil {
		h.logger.Error("failed to get users", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
