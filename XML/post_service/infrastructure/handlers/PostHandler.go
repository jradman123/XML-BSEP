package handlers

import (
	common "common/module"
	"common/module/interceptor"
	"common/module/logger"
	pb "common/module/proto/posts_service"
	"context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/infrastructure/api"
	"strings"

	"post/module/application"
)

type PostHandler struct {
	postService *application.PostService
	userService *application.UserService
	logInfo     *logger.Logger
	logError    *logger.Logger
}

func NewPostHandler(service *application.PostService, userService *application.UserService, logInfo *logger.Logger, logError *logger.Logger) *PostHandler {
	return &PostHandler{postService: service, userService: userService, logInfo: logInfo, logError: logError}
}
func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {

}

func (p PostHandler) GetAllByUsername(_ context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	//request = p.sanitizeGetRequest(request)

	posts, err := p.postService.GetAllByUsername(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET ALL POSTS FOR USER FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}

	for _, post := range posts {
		current := api.MapPostReply(post)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (p PostHandler) Get(_ context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {

	//request = p.sanitizeGetRequest(request)

	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	postPb := api.MapPostReply(post)
	response := &pb.GetResponse{Post: postPb}
	return response, nil
}

func (p PostHandler) GetAll(_ context.Context, _ *pb.Empty) (*pb.GetMultipleResponse, error) {
	p.logInfo.Logger.Infof("INFO:Handling GetAllPosts")
	posts, err := p.postService.GetAll()
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL POSTS FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}
	for _, post := range posts {
		current := api.MapPostReply(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (p PostHandler) Create(ctx context.Context, request *pb.CreatePostRequest) (*pb.Empty, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	//request = p.sanitizePost(request, userNameCtx)

	user, _ := p.userService.GetByUsername(request.Post.Username)
	post := api.MapNewPost(request.Post, user[0])
	err := p.postService.Create(post)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:CREATE POST")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	//request = p.sanitizeComment(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
	}
	comment := api.MapNewComment(request.Comment)
	err = p.postService.CreateComment(post, comment)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:CREATE COMMENT")
		return nil, err
	}

	return &pb.CreateCommentResponse{
		Comment: request.Comment,
	}, nil
}

func (p PostHandler) LikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.Empty, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	//request = p.sanitizeReactionRequest(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	fmt.Println("USERNAME")
	fmt.Println(request.Username)
	user, err := p.userService.GetByUsername(request.Username)
	fmt.Println("USER")
	fmt.Println(user[0])
	err = p.postService.LikePost(post, user[0].UserId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:LIKE POST")
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (p PostHandler) DislikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.Empty, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	//request = p.sanitizeReactionRequest(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	user, err := p.userService.GetByUsername(request.Username)
	err = p.postService.DislikePost(post, user[0].UserId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:DISLIKE POST")
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (p PostHandler) CreateJobOffer(_ context.Context, request *pb.CreateJobOfferRequest) (*pb.Empty, error) {
	request = p.sanitizeJobOffer(request)
	offer := api.MapNewJobOffer(request.JobOffer)

	err := p.postService.CreateJobOffer(offer)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"jobOfferId": request.JobOffer.Id,
		}).Errorf("ERR:CREATE JOB OFFER")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) GetAllJobOffers(_ context.Context, _ *pb.Empty) (*pb.GetAllJobOffers, error) {
	p.logInfo.Logger.Infof("INFO:Handling GetAllJobOffers")
	offers, err := p.postService.GetAllJobOffers()
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL JOB OFFERS FROM DB")
		return nil, err
	}
	response := &pb.GetAllJobOffers{JobOffers: []*pb.JobOffer{}}
	for _, offer := range offers {
		current := api.MapJobOfferReply(offer)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (p PostHandler) GetAllReactionsForPost(_ context.Context, request *pb.GetRequest) (*pb.GetReactionsResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	sqlInj := common.BadId(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling GetAllReactionsForPost")
	}
	objectId, err := primitive.ObjectIDFromHex(request.Id)

	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}

	likesNum, dislikesNum := api.FindNumberOfReactions(post)
	response := &pb.GetReactionsResponse{}
	response.DislikesNumber = int32(dislikesNum)
	response.LikesNumber = int32(likesNum)

	return response, nil
}

func (p PostHandler) GetAllCommentsForPost(_ context.Context, request *pb.GetRequest) (*pb.GetAllCommentsResponse, error) {
	//request = p.sanitizeGetRequest(request)
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.postService.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}

	response := &pb.GetAllCommentsResponse{Comments: []*pb.Comment{}}
	for _, comment := range post.Comments {
		user, err := p.userService.GetByUsername(comment.Username)
		if err != nil {
			return nil, err
		}
		current := api.MapUserCommentsForPost(user[0], comment.CommentText)
		response.Comments = append(response.Comments, current)
	}

	return response, nil
}
