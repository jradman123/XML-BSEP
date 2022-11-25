package handlers

import (
	common "common/module"
	"common/module/interceptor"
	"common/module/logger"
	notificationProto "common/module/proto/notification_service"
	pb "common/module/proto/posts_service"
	"context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
	"post/module/infrastructure/api"
	"strings"

	tracer "monitoring/module"
	"post/module/application"
)

type PostHandler struct {
	postService        *application.PostService
	userService        *application.UserService
	logInfo            *logger.Logger
	logError           *logger.Logger
	notificationClient notificationProto.NotificationServiceClient
}

func NewPostHandler(service *application.PostService, userService *application.UserService, logInfo *logger.Logger, logError *logger.Logger) *PostHandler {
	return &PostHandler{postService: service, userService: userService, logInfo: logInfo, logError: logError}
}
func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {

}

func (p PostHandler) GetAllByUsername(ctx context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	//request = p.sanitizeGetRequest(request)
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllByUsername")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	posts, err := p.postService.GetAllByUsername(request.Id, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET ALL POSTS FOR USER FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}

	for _, post := range posts {
		current := api.MapPostReply(post, ctx)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (p PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {

	//request = p.sanitizeGetRequest(request)
	span := tracer.StartSpanFromContextMetadata(ctx, "Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	postPb := api.MapPostReply(post, ctx)
	response := &pb.GetResponse{Post: postPb}
	return response, nil
}

func (p PostHandler) GetAll(ctx context.Context, _ *pb.Empty) (*pb.GetMultipleResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	p.logInfo.Logger.Infof("INFO:Handling GetAllPosts")
	posts, err := p.postService.GetAll(ctx)
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL POSTS FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}
	for _, post := range posts {
		current := api.MapPostReply(post, ctx)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (p PostHandler) CheckLikedStatus(ctx context.Context, request *pb.UserReactionRequest) (*pb.GetUserReactionResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckLikedStatus")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		panic(err)
	}
	user, err := p.userService.GetByUsername(request.Username, ctx)
	if err != nil {
		panic(err)
	}
	reaction, err := p.postService.CheckLikedStatus(objectId, user[0].UserId, ctx)
	if err != nil {
		panic(err)
	}
	response := &pb.GetUserReactionResponse{
		Liked:    false,
		Disliked: false,
		Neutral:  false,
	}
	if reaction == model.LIKED {
		response.Liked = true
	}
	if reaction == model.DISLIKED {
		response.Disliked = true
	}
	if reaction == model.Neutral {
		response.Neutral = true
	}
	return response, nil
}

func (p PostHandler) Create(ctx context.Context, request *pb.CreatePostRequest) (*pb.Empty, error) {
	//userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	//request = p.sanitizePost(request, userNameCtx)
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, _ := p.userService.GetByUsername(request.Post.Username, ctx)
	post := api.MapNewPost(request.Post, user[0], ctx)
	err := p.postService.Create(post, ctx)

	if err != nil {
		p.logError.Logger.Errorf("ERR:CREATE POST")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	ctx = tracer.ContextWithSpan(context.Background(), span)
	//request = p.sanitizeComment(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
	}
	comment := api.MapNewComment(request.Comment, ctx)
	err = p.postService.CreateComment(post, comment, ctx)
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
	span := tracer.StartSpanFromContextMetadata(ctx, "LikePost")
	defer span.Finish()
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	ctx = tracer.ContextWithSpan(context.Background(), span)

	//request = p.sanitizeReactionRequest(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	fmt.Println("USERNAME")
	fmt.Println(request.Username)
	user, err := p.userService.GetByUsername(request.Username, ctx)
	fmt.Println("USER")
	fmt.Println(user[0])
	err = p.postService.LikePost(post, user[0].UserId, request.Username, ctx)
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
	span := tracer.StartSpanFromContextMetadata(ctx, "DislikePost")
	defer span.Finish()
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	ctx = tracer.ContextWithSpan(context.Background(), span)

	//request = p.sanitizeReactionRequest(request, userNameCtx)
	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	user, err := p.userService.GetByUsername(request.Username, ctx)
	err = p.postService.DislikePost(post, user[0].UserId, request.Username, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:DISLIKE POST")
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (p PostHandler) CreateJobOffer(ctx context.Context, request *pb.CreateJobOfferRequest) (*pb.Empty, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateJobOffer")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	request = p.sanitizeJobOffer(request)
	offer := api.MapNewJobOffer(request.JobOffer, ctx)

	err := p.postService.CreateJobOffer(offer, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"jobOfferId": request.JobOffer.Id,
		}).Errorf("ERR:CREATE JOB OFFER")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) GetAllJobOffers(ctx context.Context, _ *pb.Empty) (*pb.GetAllJobOffers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllJobOffers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	p.logInfo.Logger.Infof("INFO:Handling GetAllJobOffers")
	offers, err := p.postService.GetAllJobOffers(ctx)
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL JOB OFFERS FROM DB")
		return nil, err
	}
	response := &pb.GetAllJobOffers{JobOffers: []*pb.JobOffer{}}
	for _, offer := range offers {
		current := api.MapJobOfferReply(offer, ctx)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (p PostHandler) GetUsersJobOffers(ctx context.Context, req *pb.GetMyJobsRequest) (*pb.GetAllJobOffers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersJobOffers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	offers, err := p.postService.GetUsersJobOffers(req.Username, ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobOffers{JobOffers: []*pb.JobOffer{}}
	for _, offer := range offers {
		current := api.MapJobOfferReply(offer, ctx)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (p PostHandler) GetAllReactionsForPost(ctx context.Context, request *pb.GetRequest) (*pb.GetReactionsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllReactionsForPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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

	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}

	likesNum, dislikesNum := api.FindNumberOfReactions(post, ctx)
	response := &pb.GetReactionsResponse{}
	response.DislikesNumber = int32(dislikesNum)
	response.LikesNumber = int32(likesNum)

	return response, nil
}

func (p PostHandler) GetAllCommentsForPost(ctx context.Context, request *pb.GetRequest) (*pb.GetAllCommentsResponse, error) {
	//request = p.sanitizeGetRequest(request)

	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllCommentsForPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.postService.Get(objectId, ctx)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}

	response := &pb.GetAllCommentsResponse{Comments: []*pb.Comment{}}
	for _, comment := range post.Comments {
		user, err := p.userService.GetByUsername(comment.Username, ctx)
		if err != nil {
			return nil, err
		}
		current := api.MapUserCommentsForPost(user[0], comment.CommentText, ctx)
		response.Comments = append(response.Comments, current)
	}

	return response, nil
}
