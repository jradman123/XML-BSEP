package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
)

type PostRepository interface {
	Get(id primitive.ObjectID) (*model.Post, error)
	GetAll() ([]*model.Post, error)
	Create(post *model.Post) error
	GetAllByUserId(uuid string) ([]*model.Post, error)
	CreateComment(post *model.Post, comment *model.Comment) error
	LikePost(post *model.Post, username string) error
	DislikePost(post *model.Post, username string) error
	CreateJobOffer(offer *model.JobOffer) error
	GetAllJobOffers() ([]*model.JobOffer, error)
	UpdateUserPosts(user *model.User) error
}
