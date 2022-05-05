package handlers

import (
	"log"
	"net/http"
)

// UserHandler is a http.Handler
type UserHandler struct {
	l *log.Logger
}

// NewUserHandler creates a user handler with the given logger, simmilar to a constructor with dependency injection
func NewUserHandler(l *log.Logger) *UserHandler {
	return &UserHandler{l}
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling GET Users")
}
func (u *UserHandler) AddUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling POST Users")
}
func (u *UserHandler) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling PUT Users")
}
