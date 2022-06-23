package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	FirstName string             `bson:"name"`
	LastName  string             `bson:"surname"`
}
