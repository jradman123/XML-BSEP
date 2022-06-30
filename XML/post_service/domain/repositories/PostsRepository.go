package repositories

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
)

type PostRepository interface {
	Get(id primitive.ObjectID) (*model.Post, error)
	GetAll() ([]*model.Post, error)
	Create(post *model.Post) error
	GetAllByUsername(username string) ([]*model.Post, error)
	CreateComment(post *model.Post, comment *model.Comment) error
	LikePost(post *model.Post, userId uuid.UUID) error
	DislikePost(post *model.Post, userId uuid.UUID) error
	CreateJobOffer(offer *model.JobOffer) error
	GetAllJobOffers() ([]*model.JobOffer, error)
	UpdateUserPosts(user *model.User) error
	CheckLikedStatus(id primitive.ObjectID, userId uuid.UUID) (model.ReactionType, error)
}
