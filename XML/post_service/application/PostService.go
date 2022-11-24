package application

import (
	"common/module/logger"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	tracer "monitoring/module"
	"post/module/domain/model"
	"post/module/domain/repositories"
	"post/module/infrastructure/orchestrators"
)

type PostService struct {
	repository       repositories.PostRepository
	logInfo          *logger.Logger
	logError         *logger.Logger
	postOrchestrator *orchestrators.PostOrchestrator
	jobOrchestrator  *orchestrators.JobOrchestrator
}

func NewPostService(repository repositories.PostRepository, logInfo *logger.Logger, logError *logger.Logger, porchestrator *orchestrators.PostOrchestrator, jorchestrator *orchestrators.JobOrchestrator) *PostService {
	return &PostService{repository: repository, logInfo: logInfo, logError: logError, postOrchestrator: porchestrator, jobOrchestrator: jorchestrator}
}

func (service *PostService) Get(id primitive.ObjectID, ctx context.Context) (*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "getService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.Get(id, ctx)
}

func (service *PostService) GetAll(ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetAll(ctx)
}

func (service *PostService) Create(post *model.Post, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "createPostService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.Create(post, ctx)
}

func (service *PostService) GetAllByUsername(username string, ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllByUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetAllByUsername(username, ctx)
}

func (service *PostService) CreateComment(post *model.Post, comment *model.Comment, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "createCommentService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	service.postOrchestrator.CommentPost(post.Id, comment.Username, post.Username)
	return service.repository.CreateComment(post, comment, ctx)
}

func (service *PostService) LikePost(post *model.Post, userId uuid.UUID, likerUsername string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "likePostService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	service.postOrchestrator.LikePost(post.Id, likerUsername, post.Username)
	return service.repository.LikePost(post, userId, ctx)
}

func (service *PostService) DislikePost(post *model.Post, userId uuid.UUID, haterUsername string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "dislikePostService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	service.postOrchestrator.DislikePost(post.Id, haterUsername, post.Username)
	return service.repository.DislikePost(post, userId, ctx)
}

func (service *PostService) CreateJobOffer(offer *model.JobOffer, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "createJobOfferService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	offer, err := service.repository.CreateJobOffer(offer, ctx)
	service.jobOrchestrator.CreateJobOffer(*offer)
	return err
}

func (service *PostService) GetAllJobOffers(ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllJobOffersService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetAllJobOffers(ctx)
}

func (service *PostService) UpdateUserPosts(user *model.User, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateUserPostsService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.UpdateUserPosts(user, ctx)
}

func (service *PostService) CheckLikedStatus(id primitive.ObjectID, userId uuid.UUID, ctx context.Context) (model.ReactionType, error) {
	span := tracer.StartSpanFromContext(ctx, "checkLikedStatusService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.CheckLikedStatus(id, userId, ctx)
}

func (service *PostService) GetUsersJobOffers(username string, ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "getUsersJobOffersService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetUsersJobOffers(username, ctx)
}
