package handlers

import (
	pb "common/module/proto/posts_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/infrastructure/api"

	"post/module/application"
)

type PostHandler struct {
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{service: service}
}
func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {

}

func (p PostHandler) GetAllByUserId(ctx context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	id := request.Id
	posts, err := p.service.GetAllByUserId(id)
	if err != nil {
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
	id := request.GetId()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := api.MapPost(post)
	response := &pb.GetResponse{Post: postPb}
	return response, nil
}

func (p PostHandler) GetAll(ctx context.Context, empty *pb.Empty) (*pb.GetMultipleResponse, error) {
	posts, err := p.service.GetAll()
	if err != nil {
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
	post := api.MapNewPost(request.Post)
	err := p.service.Create(post)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	comment := api.MapNewComment(request.Comment)
	err = p.service.CreateComment(post, comment)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{
		Comment: request.Comment,
	}, nil
}

func (p PostHandler) LikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.Empty, error) {
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	err = p.service.LikePost(post, request.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (p PostHandler) DislikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.Empty, error) {
	objectId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	err = p.service.DislikePost(post, request.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (p PostHandler) CreateJobOffer(ctx context.Context, request *pb.CreateJobOfferRequest) (*pb.Empty, error) {
	offer := api.MapNewJobOffer(request.JobOffer)
	err := p.service.CreateJobOffer(offer)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (p PostHandler) GetAllJobOffers(ctx context.Context, empty *pb.Empty) (*pb.GetAllJobOffers, error) {
	offers, err := p.service.GetAllJobOffers()
	if err != nil {
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
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	post, err := p.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	likesNum, dislikesNum := api.FindNumberOfReactions(post)
	response := &pb.GetReactionsResponse{}
	response.DislikesNumber = int32(dislikesNum)
	response.LikesNumber = int32(likesNum)

	return response, nil
}

func (p PostHandler) GetAllCommentsForPost(ctx context.Context, request *pb.GetRequest) (*pb.GetAllCommentsResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	_, err = p.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllCommentsResponse{Comments: []*pb.Comment{}}
	return response, nil
}
