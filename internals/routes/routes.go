package routes

import (
	"github.com/gorilla/mux"
	"github.com/sumukhj1219/go-profiler/internals/handler"
)

func RegisterRoutes(r *mux.Router, h handler.Handler) {
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/create", h.UserHandler.CreateUser).Methods("POST")
	s.HandleFunc("/all", h.UserHandler.GetUsers).Methods("GET")
	s.HandleFunc("/email", h.UserHandler.GetUserByEmail).Methods("GET")
	s.HandleFunc("/id", h.UserHandler.GetUserById).Methods("GET")
}
