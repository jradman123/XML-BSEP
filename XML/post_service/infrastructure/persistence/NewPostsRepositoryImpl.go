package persistence

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"post/module/domain/model"
	"post/module/domain/repositories"
)

const (
	DATABASE   = "posts_service"
	COLLECTION = "postsData"
)

type PostRepositoryImpl struct {
	collection *mongo.Collection
}

func NewPostRepositoryImpl(client *mongo.Client) repositories.PostRepository {
	collection := client.Database(DATABASE).Collection(COLLECTION)
	return &PostRepositoryImpl{collection: collection}
}

func (p PostRepositoryImpl) Get(id primitive.ObjectID) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) GetAll() ([]*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) Create(post *model.Post) error {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) GetAllByUserId(uuid string) ([]*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) CreateComment(post *model.Post, comment *model.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) LikePost(post *model.Post, username string) error {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) DislikePost(post *model.Post, username string) error {
	//TODO implement me
	panic("implement me")
}
