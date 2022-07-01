package model

type Post struct {
	Id             string
	Username       string
	PostText       string
	ImagePaths     string
	DatePosted     string
	LikesNumber    int32
	DislikesNumber int32
	CommentsNumber int32
	Links          Links
}
type Links struct {
	Comment string
	Like    string
	Dislike string
	User    string
}
