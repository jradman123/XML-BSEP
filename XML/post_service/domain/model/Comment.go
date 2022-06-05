package model

type Comment struct {
	UserId      string `bson:"user_id"`
	CommentText string `bson:"comment_text"`
}
