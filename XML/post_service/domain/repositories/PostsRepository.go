package repositories

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
)

type PostRepository interface {
	Get(id primitive.ObjectID, ctx context.Context) (*model.Post, error)
	GetAll(ctx context.Context) ([]*model.Post, error)
	Create(post *model.Post, ctx context.Context) error
	GetAllByUsername(username string, ctx context.Context) ([]*model.Post, error)
	CreateComment(post *model.Post, comment *model.Comment, ctx context.Context) error
	LikePost(post *model.Post, userId uuid.UUID, ctx context.Context) error
	DislikePost(post *model.Post, userId uuid.UUID, ctx context.Context) error
	CreateJobOffer(offer *model.JobOffer, ctx context.Context) (*model.JobOffer, error)
	GetAllJobOffers(ctx context.Context) ([]*model.JobOffer, error)
	UpdateUserPosts(user *model.User, ctx context.Context) error
	CheckLikedStatus(id primitive.ObjectID, userId uuid.UUID, ctx context.Context) (model.ReactionType, error)
	GetUsersJobOffers(username string, ctx context.Context) ([]*model.JobOffer, error)
}
