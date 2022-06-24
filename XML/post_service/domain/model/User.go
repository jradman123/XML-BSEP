package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    uuid.UUID          `bson:"user_id"`
	Username  string             `bson:"username"`
	FirstName string             `bson:"name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Active    bool               `bson:"active"`
}
