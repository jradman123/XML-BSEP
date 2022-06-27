package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	CommentText string             `bson:"comment_text"`
}
