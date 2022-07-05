package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id          primitive.ObjectID
	Username    string
	CommentText string
}
