package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/post_events"
	"context"
	"message/module/application"
	"message/module/infrastructure/api"
	tracer "monitoring/module"
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

func (handler *NotificationCommandHandler) handle(command *events.PostNotificationCommand, ctx context.Context) {
	span := tracer.StartSpanFromContextMetadata(ctx, "notificationCommandHandler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	notification := api.MapNewPostNotification(command, ctx)
	var reply = &events.PostNotificationReply{}
	switch command.Type {
	case events.LikePost:
		_, err := handler.service.Create(notification, ctx)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, ctx)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, ctx)
	case events.DislikePost:
		_, err := handler.service.Create(notification, ctx)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, ctx)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, ctx)
	case events.CommentPost:
		_, err := handler.service.Create(notification, ctx)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent, ctx)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent, ctx)
	default:
		reply = api.MapPostNotificationReply(events.UnknownReply, ctx)

		if reply.Type != events.UnknownReply {
			_ = handler.replyPublisher.Publish(reply)
		}

	}

}
