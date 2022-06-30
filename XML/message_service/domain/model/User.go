package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	UserId   uuid.UUID          `bson:"user_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
}
