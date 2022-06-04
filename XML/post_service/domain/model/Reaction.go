package model

type Reaction struct {
	UserId   string       `bson:"user_id"`
	Reaction ReactionType `bson:"reaction"`
}
