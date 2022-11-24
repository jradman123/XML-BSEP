package application

import (
	"common/module/logger"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	tracer "monitoring/module"
	"post/module/domain/model"
	"post/module/domain/repositories"
)

type UserService struct {
	repository repositories.UserRepository
	logInfo    *logger.Logger
	logError   *logger.Logger
}

func NewUserService(repository repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *UserService {
	return &UserService{repository: repository, logInfo: logInfo, logError: logError}
}

func (s UserService) CreateUser(requestUser *model.User, ctx context.Context) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "createUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.CreateUser(requestUser, ctx)
	return user, err
}

func (s UserService) UpdateUser(requestUser *model.User, ctx context.Context) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "updateUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.UpdateUser(requestUser, ctx)
	return user, err
}

func (s UserService) DeleteUser(userId uuid.UUID, ctx context.Context) (err error) {
	span := tracer.StartSpanFromContext(ctx, "deleteUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	err = s.repository.DeleteUser(userId, ctx)
	return err
}

func (s UserService) ActivateUserAccount(userId uuid.UUID, ctx context.Context) (err error) {
	span := tracer.StartSpanFromContext(ctx, "activateUserAccountService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	err = s.repository.ActivateUserAccount(userId, ctx)
	return err
}
func (s UserService) Get(id primitive.ObjectID) (user *model.User, err error) {
	user, err = s.repository.Get(id)
	return user, err
}
func (s UserService) GetByUserId(id uuid.UUID) (user []*model.User, err error) {
	user, err = s.repository.GetByUserId(id)
	return user, err
}
func (s UserService) GetByUsername(username string, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "getByUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.GetByUsername(username, ctx)
	return user, err
}
