package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	UserUUID string             `bson:"userUUID"`
	Username string             `bson:"username"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
}
