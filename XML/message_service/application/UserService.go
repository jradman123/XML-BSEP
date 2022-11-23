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

func (s UserService) CreateUser(requestUser *model.User) (user *model.User, err error) {
	user, err = s.repository.CreateUser(requestUser)
	return user, err
}

func (s UserService) UpdateUser(requestUser *model.User) (user *model.User, err error) {
	user, err = s.repository.UpdateUser(requestUser)
	return user, err
}

func (s UserService) DeleteUser(userId uuid.UUID) (err error) {
	err = s.repository.DeleteUser(userId)
	return err
}
func (s UserService) GetByUsername(username string, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "getByUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.GetByUsername(username, ctx)
	return user, err
}

func (s UserService) GetSettingsForUser(username string, ctx context.Context) (settings *model.NotificationSettings, err error) {
	span := tracer.StartSpanFromContext(ctx, "getSettingsForUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.repository.GetSettingsForUser(username, ctx)
}

func (s UserService) ChangeSettingsForUser(username string, newSettings *model.NotificationSettings, ctx context.Context) (settings *model.NotificationSettings, err error) {
	span := tracer.StartSpanFromContext(ctx, "changeSettingsForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.repository.ChangeSettingsForUser(username, newSettings, ctx)
}

func (s UserService) AllowedNotificationForUser(username string, notificationType model.NotificationType, ctx context.Context) (result bool, err error) {
	span := tracer.StartSpanFromContext(ctx, "allowedNotificationForUser")
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
	span := tracer.StartSpanFromContext(ctx, "getByIdService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err = s.repository.GetById(userId, ctx)
	return user, err
}
