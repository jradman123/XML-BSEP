package handlers

import (
	"common/module/logger"
	connectionPb "common/module/proto/connection_service"
	postPb "common/module/proto/posts_service"
	"context"
	"encoding/json"
	"fmt"
	"gateway/module/domain/dto"
	"gateway/module/domain/model"
	clients "gateway/module/infrastructure/api"
	"gateway/module/startup/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	otgo "github.com/opentracing/opentracing-go"
	tracer "monitoring/module"
	"net/http"
)

type UserFeedHandler struct {
	logInfo                  *logger.Logger
	logError                 *logger.Logger
	config                   *config.Config
	userServiceAddress       string
	postServiceAddress       string
	connectionServiceAddress string
}

func NewUserFeedHandler(logInfo *logger.Logger, logError *logger.Logger, c *config.Config) Handler {
	return &UserFeedHandler{
		logInfo:                  logInfo,
		logError:                 logError,
		config:                   c,
		userServiceAddress:       fmt.Sprintf("%s:%s", c.UserHost, c.UserPort),
		postServiceAddress:       fmt.Sprintf("%s:%s", c.PostsHost, c.PostsPort),
		connectionServiceAddress: fmt.Sprintf("%s:%s", c.ConnectionsHost, c.ConnectionsPort),
	}
}

func (u UserFeedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/users/{username}/feed", u.GetFeedPostsForUser)
	if err != nil {
		panic(err)
	}

}

func (u UserFeedHandler) GetFeedPostsForUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	span := tracer.StartSpanFromRequest("GetFeedPostsForUser", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)
	fmt.Println("GetFeedPostsForUser HANDLER")
	username := params["username"]
	if username == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	postsClient := clients.NewPostClient(u.postServiceAddress)
	connectionClient := clients.NewConnectionClient(u.connectionServiceAddress)

	connections, err := connectionClient.GetConnections(ctx, &connectionPb.GetRequest{Username: username})
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var posts []model.Post
	for _, user := range connections.Users {
		postsPbs, err := postsClient.GetAllByUsername(ctx, &postPb.GetRequest{Id: user.Username})
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		posts = append(posts, mapPbsToModel(postsPbs.Posts)...)
	}
	fmt.Println("duzina svih postova")
	fmt.Println(len(posts))

	postsDto := dto.FeedPostsResponseDto{
		Feed: posts,
	}

	postsDtoJson, _ := json.Marshal(postsDto)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(postsDtoJson)
	if err != nil {
		return
	}
}
func mapPbsToModel(postsPb []*postPb.Post) []model.Post {
	var posts []model.Post

	for _, postPb := range postsPb {
		var post model.Post
		post.Id = postPb.Id
		post.Username = postPb.Username
		post.PostText = postPb.PostText
		post.ImagePaths = postPb.ImagePaths
		post.DatePosted = postPb.DatePosted
		post.LikesNumber = postPb.LikesNumber
		post.DislikesNumber = postPb.DislikesNumber
		post.CommentsNumber = postPb.CommentsNumber
		post.Links = model.Links{Comment: postPb.Links.Comment, Dislike: postPb.Links.Dislike, Like: postPb.Links.Like}

		posts = append(posts, post)
	}
	return posts
}
