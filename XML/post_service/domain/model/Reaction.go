package model

type Reaction struct {
	Username string       `bson:"username"`
	Reaction ReactionType `bson:"reaction"`
}
