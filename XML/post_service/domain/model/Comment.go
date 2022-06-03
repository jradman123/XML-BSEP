package model

type Comment struct {
	Username    string `bson:"username"`
	CommentText string `bson:"comment_text"`
}
