package handler

import (
	"net/http"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByEmail(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	UserHandler UserHandler
}
