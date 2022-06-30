package model

type Reaction struct {
	UserId   string       `bson:"user_id"`
	Reaction ReactionType `bson:"reaction"`
}

type ReactionType int

const (
	Neutral ReactionType = iota
	LIKED
	DISLIKED
)
