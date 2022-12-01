package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/post_events"
	"context"
	"message/module/application"
	"message/module/infrastructure/api"
)

type NotificationCommandHandler struct {
	service           *application.NotificationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewNotificationCommandHandler(service *application.NotificationService, publisher saga.Publisher, subscriber saga.Subscriber) (*NotificationCommandHandler, error) {
	o := &NotificationCommandHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *NotificationCommandHandler) handle(command *events.PostNotificationCommand) {

	notification := api.MapNewPostNotification(command, context.TODO())
	var reply = &events.PostNotificationReply{}
	switch command.Type {
	case events.LikePost:
		_, err := handler.service.Create(notification, context.TODO())
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, context.TODO())
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, context.TODO())
	case events.DislikePost:
		_, err := handler.service.Create(notification, context.TODO())
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, context.TODO())
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, context.TODO())
	case events.CommentPost:
		_, err := handler.service.Create(notification, context.TODO())
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, context.TODO())
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, context.TODO())
	default:
		reply = api.MapPostNotificationReply(events.UnknownReply, context.TODO())

		if reply.Type != events.UnknownReply {
			_ = handler.replyPublisher.Publish(reply)
		}

	}

}
