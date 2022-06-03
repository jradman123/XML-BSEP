package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
	"post/module/domain/repositories"
)

type PostService struct {
	repository repositories.PostRepository
}

func NewPostService(repository repositories.PostRepository) *PostService {
	return &PostService{repository: repository}
}

func (service *PostService) Get(id primitive.ObjectID) (*model.Post, error) {
	return service.repository.Get(id)
}

func (service *PostService) GetAll() ([]*model.Post, error) {
	return service.repository.GetAll()
}

func (service *PostService) Create(post *model.Post) error {
	return service.repository.Create(post)
}

func (service *PostService) GetAllByUserId(uuid string) ([]*model.Post, error) {
	return service.repository.GetAllByUserId(uuid)
}

func (service *PostService) CreateComment(post *model.Post, comment *model.Comment) error {
	return service.repository.CreateComment(post, comment)
}

func (service *PostService) LikePost(post *model.Post, username string) error {
	return service.repository.LikePost(post, username)
}

func (service *PostService) DislikePost(post *model.Post, username string) error {
	return service.repository.DislikePost(post, username)
}
