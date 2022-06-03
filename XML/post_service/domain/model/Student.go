package model

type Student struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Age  string `bson:"age"`
}
