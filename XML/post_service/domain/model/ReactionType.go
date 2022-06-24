package model

type ReactionType int

const (
	Neutral ReactionType = iota
	LIKED
	DISLIKED
)
