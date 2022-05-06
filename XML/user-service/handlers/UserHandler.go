package handlers

import (
	"log"
	"mime"
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

func (u *UserHandler) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling PUT Users")
}

func (u *UserHandler) AddUsers(w http.ResponseWriter, req *http.Request) {
	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(nil, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//id := ts.store.CreatePost(ctx, rt.Title, rt.Text, rt.Tags)
	//renderJSON(ctx, w, ResponseId{Id: id})
}
