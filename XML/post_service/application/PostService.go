package application

import (
	"common/module/logger"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
	"post/module/domain/repositories"
	"post/module/infrastructure/orchestrators"
)

type PostService struct {
	repository   repositories.PostRepository
	logInfo      *logger.Logger
	logError     *logger.Logger
	orchestrator *orchestrators.JobOrchestrator
}

func NewPostService(repository repositories.PostRepository, logInfo *logger.Logger, logError *logger.Logger, orchestrator *orchestrators.JobOrchestrator) *PostService {
	return &PostService{repository: repository, logInfo: logInfo, logError: logError, orchestrator: orchestrator}
}

func (service *PostService) Get(id primitive.ObjectID) (*model.Post, error) {
	return service.repository.Get(id)
}

func (service *PostService) GetAll() ([]*model.Post, error) {
	return service.repository.GetAll()
}

func (service *PostService) Create(post *model.Post) error {
	return service.repository.Create(post)
}

func (service *PostService) GetAllByUsername(username string) ([]*model.Post, error) {
	return service.repository.GetAllByUsername(username)
}

func (service *PostService) CreateComment(post *model.Post, comment *model.Comment) error {
	return service.repository.CreateComment(post, comment)
}

func (service *PostService) LikePost(post *model.Post, userId uuid.UUID) error {
	return service.repository.LikePost(post, userId)
}

func (service *PostService) DislikePost(post *model.Post, userId uuid.UUID) error {
	return service.repository.DislikePost(post, userId)
}

func (service *PostService) CreateJobOffer(offer *model.JobOffer) error {
	offer, err := service.repository.CreateJobOffer(offer)
	service.orchestrator.CreateJobOffer(*offer)
	return err
}

func (service *PostService) GetAllJobOffers() ([]*model.JobOffer, error) {
	return service.repository.GetAllJobOffers()
}

func (service *PostService) UpdateUserPosts(user *model.User) error {
	return service.repository.UpdateUserPosts(user)
}

func (service *PostService) CheckLikedStatus(id primitive.ObjectID, userId uuid.UUID) (model.ReactionType, error) {
	return service.repository.CheckLikedStatus(id, userId)
}

func (service *PostService) GetUsersJobOffers(username string) ([]*model.JobOffer, error) {
	return service.repository.GetUsersJobOffers(username)
}
