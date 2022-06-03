package handlers

import "post/module/application"

type PostHandler struct {
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{service: service}
}
