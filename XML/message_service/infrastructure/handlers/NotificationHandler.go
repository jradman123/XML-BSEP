package handlers

import (
	"common/module/logger"
	notificationProto "common/module/proto/notification_service"
	pb "common/module/proto/notification_service"
	"context"
	"github.com/pusher/pusher-http-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/application"
	"message/module/domain/model"
	"message/module/infrastructure/api"
	"time"
)

type NotificationHandler struct {
	logInfo             *logger.Logger
	logError            *logger.Logger
	notificationPusher  *pusher.Client
	notificationService *application.NotificationService
	userService         *application.UserService
}

func (n NotificationHandler) GetAllSent(ctx context.Context, request *notificationProto.GetRequest) (*notificationProto.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationHandler) MustEmbedUnimplementedNotificationServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewNotificationHandler(logInfo *logger.Logger, logError *logger.Logger, notificationPusher *pusher.Client, notificationService *application.NotificationService, userService *application.UserService) *NotificationHandler {
	return &NotificationHandler{logInfo, logError, notificationPusher, notificationService, userService}
}

func (n NotificationHandler) Create(ctx context.Context, newNotificationReq *pb.NewNotificationRequest) (*pb.Empty, error) {
	// create Notification object and store it in the database
	// trigger pusher
	// check if this notification is blocked for that user

	notiType := model.PROFILE
	if newNotificationReq.NewNotification.NotificationType == "MESSAGE" {
		notiType = model.MESSAGE
	} else if newNotificationReq.NewNotification.NotificationType == "POST" {
		notiType = model.POST
	}

	noti := &model.Notification{
		Id:               primitive.NewObjectID(),
		Timestamp:        time.Now(),
		Content:          newNotificationReq.NewNotification.Content,
		NotificationFrom: newNotificationReq.NewNotification.From,
		NotificationTo:   newNotificationReq.NewNotification.To,
		Read:             false,
		RedirectPath:     newNotificationReq.NewNotification.RedirectPath,
		Type:             model.NotificationType(notiType),
	}

	notification := n.notificationService.Create(noti)

	n.notificationPusher.Trigger("notifications", "notification", notification)

	return &pb.Empty{}, nil
}

func (n NotificationHandler) GetAllForUser(_ context.Context, request *pb.GetAllNotificationRequest) (*pb.GetAllNotificationResponse, error) {

	notifications, _ := n.notificationService.GetAllForUser(request.Username)
	response := &pb.GetAllNotificationResponse{Notifications: []*pb.Notification{}}

	for _, notification := range notifications {
		current := api.MapNotificationResponse(notification)
		response.Notifications = append(response.Notifications, current)
	}

	return response, nil
}

func (n NotificationHandler) GetSettingsForUser(_ context.Context, request *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	settings, _ := n.userService.GetSettingsForUser(request.Username)
	response := &pb.GetSettingsResponse{Settings: &pb.NotificationSettings{}}

	response.Settings = api.MapSettingsResponse(settings)

	return response, nil
}

func (n NotificationHandler) ChangeSettingsForUser(_ context.Context, request *pb.ChangeSettingsRequest) (*pb.GetSettingsResponse, error) {

	settingsMapped := api.MapChangeSettingsRequest(request)
	settings, err := n.userService.ChangeSettingsForUser(request.NewSettings.Username, settingsMapped)
	if err != nil {
		return nil, err
	}
	response := &pb.GetSettingsResponse{Settings: &pb.NotificationSettings{}}
	response.Settings = api.MapSettingsResponse(settings)

	return response, nil
}
