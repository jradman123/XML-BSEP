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
	service  *application.PostService
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewPostHandler(service *application.PostService, logInfo *logger.Logger, logError *logger.Logger) *PostHandler {
	return &PostHandler{service: service, logInfo: logInfo, logError: logError}
}
func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {

}

func (p PostHandler) GetAllByUserId(ctx context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	sqlInj := common.BadId(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling GetAllByUserId posts")
	}
	id := request.Id
	posts, err := p.service.GetAllByUserId(id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET ALL POSTS FOR USER FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}

	for _, post := range posts {
		current := api.MapPost(post)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (p PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	sqlInj := common.BadId(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling Get post")
	}

	id := request.GetId()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	postPb := api.MapPost(post)
	response := &pb.GetResponse{Post: postPb}
	return response, nil
}

func (p PostHandler) GetAll(ctx context.Context, empty *pb.Empty) (*pb.GetMultipleResponse, error) {
	p.logInfo.Logger.Infof("INFO:Handling GetAllPosts")
	posts, err := p.service.GetAll()
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL POSTS FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}
	for _, post := range posts {
		current := api.MapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (p PostHandler) Create(ctx context.Context, request *pb.CreatePostRequest) (*pb.Empty, error) {
	policy := bluemonday.UGCPolicy()
	request.Post.UserId = strings.TrimSpace(policy.Sanitize(request.Post.UserId))
	request.Post.PostText = strings.TrimSpace(policy.Sanitize(request.Post.PostText))
	for i, _ := range request.Post.ImagePaths {
		request.Post.ImagePaths[i] = strings.TrimSpace(policy.Sanitize(request.Post.ImagePaths[i]))
	}
	request.Post.DatePosted = strings.TrimSpace(policy.Sanitize(request.Post.DatePosted))

	p1 := common.BadId(request.Post.UserId)
	p2 := common.BadText(request.Post.PostText)
	p3 := common.BadDate(request.Post.DatePosted)
	p4 := common.BadPaths(request.Post.ImagePaths)
	//sqlInj := common.CheckForSQLInjection([]string{request.Post.PostText, request.Post.UserId, request.Post.DatePosted})
	//sqlInj2 := common.CheckForSQLInjection(request.Post.ImagePaths)
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.Post.UserId == "" || request.Post.PostText == "" || request.Post.DatePosted == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 || p3 || p4 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling Create post")
	}
	post := api.MapNewPost(request.Post)
	err := p.service.Create(post)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:CREATE POST")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.PostId = strings.TrimSpace(policy.Sanitize(request.PostId))
	request.Comment.UserId = strings.TrimSpace(policy.Sanitize(request.Comment.UserId))
	request.Comment.Username = strings.TrimSpace(policy.Sanitize(request.Comment.Username))
	request.Comment.Name = strings.TrimSpace(policy.Sanitize(request.Comment.Name))
	request.Comment.Surname = strings.TrimSpace(policy.Sanitize(request.Comment.Surname))
	request.Comment.CommentText = strings.TrimSpace(policy.Sanitize(request.Comment.CommentText))

	p1 := common.BadId(request.PostId)
	p2 := common.BadId(request.Comment.UserId)
	p3 := common.BadUsername(request.Comment.Username)
	p4 := common.BadName(request.Comment.Name)
	p5 := common.BadName(request.Comment.Surname)
	p6 := common.BadText(request.Comment.CommentText)

	//sqlInj := common.CheckForSQLInjection([]string{request.PostId, request.Comment.UserId, request.Comment.Username,
	//	request.Comment.Name, request.Comment.Surname, request.Comment.CommentText})
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.PostId == "" || request.Comment.UserId == "" || request.Comment.Username == "" || request.Comment.Name == "" ||
		request.Comment.Surname == "" || request.Comment.CommentText == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 || p3 || p4 || p5 || p6 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling CreateComment")
	}
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	comment := api.MapNewComment(request.Comment)
	err = p.service.CreateComment(post, comment)
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
	policy := bluemonday.UGCPolicy()
	request.PostId = strings.TrimSpace(policy.Sanitize(request.PostId))
	request.UserId = strings.TrimSpace(policy.Sanitize(request.UserId))
	p1 := common.BadId(request.PostId)
	p2 := common.BadId(request.UserId)
	//sqlInj := common.CheckForSQLInjection([]string{request.PostId, request.UserId})
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.PostId == "" || request.UserId == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling LikePost")
	}
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	err = p.service.LikePost(post, request.UserId)
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
	policy := bluemonday.UGCPolicy()
	request.PostId = strings.TrimSpace(policy.Sanitize(request.PostId))
	request.UserId = strings.TrimSpace(policy.Sanitize(request.UserId))
	p1 := common.BadId(request.PostId)
	p2 := common.BadId(request.UserId)
	//sqlInj := common.CheckForSQLInjection([]string{request.PostId, request.UserId})
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.PostId == "" || request.UserId == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling DislikePost")
	}
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"user":   userNameCtx,
			"postId": request.PostId,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}
	err = p.service.DislikePost(post, request.UserId)
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
	policy := bluemonday.UGCPolicy()
	request.JobOffer.Publisher = strings.TrimSpace(policy.Sanitize(request.JobOffer.Publisher))
	request.JobOffer.Position = strings.TrimSpace(policy.Sanitize(request.JobOffer.Position))
	request.JobOffer.JobDescription = strings.TrimSpace(policy.Sanitize(request.JobOffer.JobDescription))
	for i, _ := range request.JobOffer.Requirements {
		request.JobOffer.Requirements[i] = strings.TrimSpace(policy.Sanitize(request.JobOffer.Requirements[i]))
	}
	request.JobOffer.DatePosted = strings.TrimSpace(policy.Sanitize(request.JobOffer.DatePosted))
	request.JobOffer.Duration = strings.TrimSpace(policy.Sanitize(request.JobOffer.Duration))
	//p1 := common.BadJWTToken(request.ShareJobOffer.ApiToken)
	p2 := common.BadText(request.JobOffer.Publisher)
	p3 := common.BadText(request.JobOffer.Position)
	p4 := common.BadText(request.JobOffer.JobDescription)
	p5 := common.BadTexts(request.JobOffer.Requirements)
	p6 := common.BadDate(request.JobOffer.DatePosted)
	p7 := common.BadDate(request.JobOffer.Duration)
	//sqlInj := common.CheckForSQLInjection([]string{request.JobOffer.Publisher,
	//	request.JobOffer.Position, request.JobOffer.JobDescription, request.JobOffer.DatePosted,
	//	request.JobOffer.Duration})
	//sqlInj2 := common.CheckForSQLInjection(request.JobOffer.Requirements)

	if request.JobOffer.Publisher == "" || request.JobOffer.Position == "" || request.JobOffer.JobDescription == "" ||
		request.JobOffer.DatePosted == "" || request.JobOffer.Duration == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if p2 || p3 || p4 || p5 || p6 || p7 {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling ShareJobOffer/CreateJobOffer")
	}
	offer := api.MapNewJobOffer(request.JobOffer)
	err := p.service.CreateJobOffer(offer)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"jobOfferId": request.JobOffer.Id,
		}).Errorf("ERR:CREATE JOB OFFER")
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) GetAllJobOffers(ctx context.Context, empty *pb.Empty) (*pb.GetAllJobOffers, error) {
	p.logInfo.Logger.Infof("INFO:Handling GetAllJobOffers")
	offers, err := p.service.GetAllJobOffers()
	if err != nil {
		p.logError.Logger.Errorf("ERR:GETTING ALL JOB OFFERS FROM DB")
		return nil, err
	}
	response := &pb.GetAllJobOffers{JobOffers: []*pb.JobOffer{}}
	for _, offer := range offers {
		current := api.MapJobOffer(offer)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (p PostHandler) GetAllReactionsForPost(ctx context.Context, request *pb.GetRequest) (*pb.GetReactionsResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	sqlInj := common.BadId(request.Id)
	//sqlInj := common.CheckRegexSQL(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling GetAllReactionsForPost")
	}
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	post, err := p.service.Get(objectId)
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

func (p PostHandler) GetAllCommentsForPost(ctx context.Context, request *pb.GetRequest) (*pb.GetAllCommentsResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	//sqlInj := common.CheckRegexSQL(request.Id)
	sqlInj := common.BadId(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling GetAllCommentsForPost")
	}
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"postId": request.Id,
		}).Errorf("ERR:HEX STRING INVALID")
		return nil, err
	}
	_, err = p.service.Get(objectId)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET POST FROM DB")
		return nil, err
	}

	response := &pb.GetAllCommentsResponse{Comments: []*pb.Comment{}}
	return response, nil
}
