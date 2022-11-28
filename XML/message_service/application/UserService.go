package application

import (
	"common/module/logger"
	"context"
	"github.com/google/uuid"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
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
	span := tracer.StartSpanFromContext(ctx, "CreateUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.CreateUser(requestUser, ctx)
	return user, err
}

func (s UserService) UpdateUser(requestUser *model.User, ctx context.Context) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "UpdateUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.UpdateUser(requestUser, ctx)
	return user, err
}

func (s UserService) DeleteUser(userId uuid.UUID, ctx context.Context) (err error) {
	span := tracer.StartSpanFromContext(ctx, "DeleteUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	err = s.repository.DeleteUser(userId, ctx)
	return err
}
func (s UserService) GetByUsername(username string, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetByUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.GetByUsername(username, ctx)
	return user, err
}

func (s UserService) GetSettingsForUser(username string, ctx context.Context) (settings *model.NotificationSettings, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetSettingsForUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.repository.GetSettingsForUser(username, ctx)
}

func (s UserService) ChangeSettingsForUser(username string, newSettings *model.NotificationSettings, ctx context.Context) (settings *model.NotificationSettings, err error) {
	span := tracer.StartSpanFromContext(ctx, "ChangeSettingsForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.repository.ChangeSettingsForUser(username, newSettings, ctx)
}

func (s UserService) AllowedNotificationForUser(username string, notificationType model.NotificationType, ctx context.Context) (result bool, err error) {
	span := tracer.StartSpanFromContext(ctx, "AllowedNotificationForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	settings, err := s.repository.GetSettingsForUser(username, ctx)
	switch notificationType {
	case model.PROFILE:
		return settings.Connections, nil
	case model.POST:
		return settings.Posts, nil
	case model.MESSAGE:
		return settings.Messages, nil
	default:
		return false, nil

	}
}
func (s UserService) GetById(userId uuid.UUID, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetByIdService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.GetById(userId, ctx)
	return user, err
}
