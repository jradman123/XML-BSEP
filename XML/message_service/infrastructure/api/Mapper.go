package api

import (
	pb "common/module/proto/message_service"
	notificationPb "common/module/proto/notification_service"
	connectionEvents "common/module/saga/connection_events"
	postEvents "common/module/saga/post_events"
	events "common/module/saga/user_events"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/domain/model"
	tracer "monitoring/module"
	"time"
)

func MapNewUser(command *events.UserCommand, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapNewUser")
	defer span.Finish()
	user := &model.User{
		Id:       primitive.NewObjectID(),
		UserId:   command.User.UserId,
		Username: command.User.Username,
		Email:    command.User.Email,
		Settings: model.NotificationSettings{
			Posts:       true,
			Messages:    true,
			Connections: true,
		},
	}
	return user
}

func MapUserForUpdate(command *events.UserCommand, userForUpdate *model.User) *model.User {
	user := &model.User{
		Id:       primitive.NewObjectID(),
		UserId:   command.User.UserId,
		Username: command.User.Username,
		Email:    command.User.Email,
		Settings: model.NotificationSettings{
			Posts:       userForUpdate.Settings.Posts,
			Messages:    userForUpdate.Settings.Messages,
			Connections: userForUpdate.Settings.Connections,
		},
	}
	return user
}

func MapNewPostNotification(command *postEvents.PostNotificationCommand, ctx context.Context) *model.Notification {
	span := tracer.StartSpanFromContext(ctx, "mapNewPostNotification")
	defer span.Finish()
	notification := &model.Notification{
		Id:               primitive.NewObjectID(),
		Timestamp:        time.Now(),
		Content:          command.Notification.Content,
		RedirectPath:     command.Notification.RedirectPath,
		Read:             false,
		Type:             model.POST,
		NotificationFrom: command.Notification.NotificationFrom,
		NotificationTo:   command.Notification.NotificationTo,
	}

	return notification
}

func MapNewConnectionNotification(command *connectionEvents.ConnectionNotificationCommand, ctx context.Context) *model.Notification {
	span := tracer.StartSpanFromContext(ctx, "mapNewConnectionNotification")
	defer span.Finish()
	notification := &model.Notification{
		Id:               primitive.NewObjectID(),
		Timestamp:        time.Now(),
		Content:          command.Notification.Content,
		RedirectPath:     command.Notification.RedirectPath,
		Read:             false,
		Type:             model.PROFILE,
		NotificationFrom: command.Notification.NotificationFrom,
		NotificationTo:   command.Notification.NotificationTo,
	}

	return notification
}

func MapPostNotificationReply(replyType postEvents.PostNotificationReplyType, ctx context.Context) (reply *postEvents.PostNotificationReply) {
	span := tracer.StartSpanFromContext(ctx, "mapPostNotificationReply")
	defer span.Finish()
	reply = &postEvents.PostNotificationReply{
		Type: replyType,
	}
	return reply
}
func MapConnectionNotificationReply(replyType connectionEvents.ConnectionNotificationReplyType, ctx context.Context) (reply *connectionEvents.ConnectionNotificationReply) {
	span := tracer.StartSpanFromContext(ctx, "mapConnectionNotificationReply")
	defer span.Finish()
	reply = &connectionEvents.ConnectionNotificationReply{
		Type: replyType,
	}
	return reply
}

func MapUserReply(user *model.User, replyType events.UserReplyType, ctx context.Context) (reply *events.UserReply) {
	span := tracer.StartSpanFromContext(ctx, "mapUserReply")
	defer span.Finish()

	reply = &events.UserReply{
		Type: replyType,
		PostUser: events.PostUser{
			Id:       user.Id,
			UserId:   user.UserId,
			Username: user.Username,
			Email:    user.Email,
		},
	}
	return reply
}

func MapMessageReply(message *model.Message, receiver string, sender string, ctx context.Context) (reply *pb.Message) {
	span := tracer.StartSpanFromContext(ctx, "mapMessageReply")
	defer span.Finish()

	reply = &pb.Message{
		SenderUsername:   sender,
		ReceiverUsername: receiver,
		MessageText:      message.MessageText,
		TimeSent:         message.TimeSent.String(),
	}
	return reply
}
func MapNewMessage(messageText string, receiverId uuid.UUID, senderId uuid.UUID, ctx context.Context) *model.Message {
	span := tracer.StartSpanFromContext(ctx, "mapNewMessage")
	defer span.Finish()

	message := &model.Message{
		Id:          primitive.NewObjectID(),
		SenderId:    senderId,
		ReceiverId:  receiverId,
		MessageText: messageText,
		TimeSent:    time.Now(),
	}
	return message
}

func MapNotificationResponse(notification *model.Notification, ctx context.Context) *notificationPb.Notification {
	span := tracer.StartSpanFromContext(ctx, "mapNotificationResponse")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := notification.Id.Hex()

	notificationPb := &notificationPb.Notification{
		Id:               id,
		Content:          notification.Content,
		From:             notification.NotificationFrom,
		To:               notification.NotificationTo,
		RedirectPath:     notification.RedirectPath,
		NotificationType: mapNotificationTypeToString(notification.Type, ctx),
		Read:             notification.Read,
		Time:             notification.Timestamp.String(),
	}

	return notificationPb
}

func MapSettingsResponse(settings *model.NotificationSettings, ctx context.Context) *notificationPb.NotificationSettings {
	span := tracer.StartSpanFromContext(ctx, "mapSettingsResponse")
	defer span.Finish()

	settingsPb := &notificationPb.NotificationSettings{
		Posts:       settings.Posts,
		Messages:    settings.Messages,
		Connections: settings.Connections,
	}

	return settingsPb
}

func MapChangeSettingsRequest(request *notificationPb.ChangeSettingsRequest, ctx context.Context) *model.NotificationSettings {
	span := tracer.StartSpanFromContext(ctx, "mapChangeSettingsRequest")
	defer span.Finish()

	settingsModel := &model.NotificationSettings{
		Posts:       request.Settings.Posts,
		Messages:    request.Settings.Messages,
		Connections: request.Settings.Connections,
	}

	return settingsModel
}

func mapNotificationTypeToString(notificationType model.NotificationType, ctx context.Context) string {
	span := tracer.StartSpanFromContext(ctx, "mapNotificationTypeToString")
	defer span.Finish()
	if notificationType == model.POST {
		return "POST"
	}
	if notificationType == model.MESSAGE {
		return "MESSAGE"
	} else {
		return "PROFILE"
	}

}
