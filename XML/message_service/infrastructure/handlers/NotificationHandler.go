package handlers

import (
	"common/module/logger"
	pb "common/module/proto/notification_service"
	"context"
	"fmt"
	"github.com/pusher/pusher-http-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/application"
	"message/module/domain/model"
	"message/module/infrastructure/api"
	tracer "monitoring/module"
	"time"
)

type NotificationHandler struct {
	logInfo             *logger.Logger
	logError            *logger.Logger
	notificationPusher  *pusher.Client
	notificationService *application.NotificationService
	userService         *application.UserService
}

func (n NotificationHandler) MustEmbedUnimplementedNotificationServiceServer() {
}

func NewNotificationHandler(logInfo *logger.Logger, logError *logger.Logger, notificationPusher *pusher.Client, notificationService *application.NotificationService, userService *application.UserService) *NotificationHandler {
	return &NotificationHandler{logInfo, logError, notificationPusher, notificationService, userService}
}

func (n NotificationHandler) Create(ctx context.Context, newNotificationReq *pb.NewNotificationRequest) (*pb.NewNotificationResponse, error) {
	// create Notification object and store it in the database
	// trigger pusher
	// check if this notification is blocked for that user
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification-Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notiType := model.PROFILE
	if newNotificationReq.NewNotification.NotificationType == "MESSAGE" {
		notiType = model.MESSAGE
	} else if newNotificationReq.NewNotification.NotificationType == "POST" {
		notiType = model.POST
	}

	result, _ := n.userService.AllowedNotificationForUser(newNotificationReq.NewNotification.To, notiType, ctx)
	if result {
		fmt.Println("dobija obavjestenja")
		noti := &model.Notification{
			Id:               primitive.NewObjectID(),
			Timestamp:        time.Now(),
			Content:          newNotificationReq.NewNotification.Content,
			NotificationFrom: newNotificationReq.NewNotification.From,
			NotificationTo:   newNotificationReq.NewNotification.To,
			Read:             false,
			RedirectPath:     newNotificationReq.NewNotification.RedirectPath,
			Type:             notiType,
		}

		notification, _ := n.notificationService.Create(noti, ctx)

		response := &pb.NewNotificationResponse{Notification: &pb.Notification{}}
		response.Notification = api.MapNotificationResponse(notification)
		return response, nil
	}
	return &pb.NewNotificationResponse{}, nil
}

func (n NotificationHandler) GetAllForUser(ctx context.Context, request *pb.GetAllNotificationRequest) (*pb.GetAllNotificationResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllForUser-Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	span1 := tracer.StartSpanFromContext(ctx, "ReadAllForUser")
	notifications, _ := n.notificationService.GetAllForUser(request.Username)
	span1.Finish()

	response := &pb.GetAllNotificationResponse{Notifications: []*pb.Notification{}}

	for _, notification := range notifications {
		current := api.MapNotificationResponse(notification)
		response.Notifications = append(response.Notifications, current)
	}

	return response, nil
}

func (n NotificationHandler) GetSettingsForUser(ctx context.Context, request *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetSettingsForUser-Handler")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	span1 := tracer.StartSpanFromContext(ctx, "ReadSettingsForUser")
	settings, _ := n.userService.GetSettingsForUser(request.Username)
	span1.Finish()

	response := &pb.GetSettingsResponse{Settings: &pb.NotificationSettings{}}

	response.Settings = api.MapSettingsResponse(settings)

	return response, nil
}

func (n NotificationHandler) ChangeSettingsForUser(ctx context.Context, request *pb.ChangeSettingsRequest) (*pb.GetSettingsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ChangeSettingsForUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	fmt.Println("usao u change settings")

	fmt.Println(*request.Settings)

	settingsMapped := api.MapChangeSettingsRequest(request)
	span1 := tracer.StartSpanFromContext(ctx, "WriteChangedSettingsForUser")
	settings, err := n.userService.ChangeSettingsForUser(request.Username, settingsMapped)
	span1.Finish()

	if err != nil {
		tracer.LogError(span1, err)
		fmt.Println(err)
		return &pb.GetSettingsResponse{}, err
	}
	response := &pb.GetSettingsResponse{Settings: &pb.NotificationSettings{}}
	response.Settings = api.MapSettingsResponse(settings)

	return response, nil
}

func (n NotificationHandler) MarkAsRead(ctx context.Context, request *pb.MarkAsReadRequest) (*pb.Empty, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "MarkAsRead")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	objectId, _ := primitive.ObjectIDFromHex(request.Id)

	span1 := tracer.StartSpanFromContext(ctx, "MarkAsReadInDatabase")
	n.notificationService.MarkAsRead(objectId)
	span1.Finish()

	return &pb.Empty{}, nil
}
