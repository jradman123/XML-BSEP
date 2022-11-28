package handlers

import (
	events "common/module/saga/connection_events"
	saga "common/module/saga/messaging"
	"context"
	"message/module/application"
	"message/module/infrastructure/api"
	tracer "monitoring/module"
)

type ConnectionNotificationCommandHandler struct {
	service           *application.NotificationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewConnectionNotificationCommandHandler(service *application.NotificationService, publisher saga.Publisher, subscriber saga.Subscriber) (*ConnectionNotificationCommandHandler, error) {
	o := &ConnectionNotificationCommandHandler{
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

func (handler *ConnectionNotificationCommandHandler) handle(command *events.ConnectionNotificationCommand, ctx context.Context) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ConnectionNotificationHandler")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	notification := api.MapNewConnectionNotification(command, ctx)
	var reply = &events.ConnectionNotificationReply{}
	switch command.Type {
	case events.Connect:
		_, err := handler.service.Create(notification, ctx)
		if err != nil {
			reply = api.MapConnectionNotificationReply(events.NotificationNotSent, ctx)
		}
		reply = api.MapConnectionNotificationReply(events.NotificationSent, ctx)
	case events.AcceptRequest:
		_, err := handler.service.Create(notification, ctx)
		if err != nil {
			reply = api.MapConnectionNotificationReply(events.NotificationNotSent, ctx)
		}
		reply = api.MapConnectionNotificationReply(events.NotificationSent, ctx)
	default:
		reply = api.MapConnectionNotificationReply(events.UnknownReply, ctx)

		if reply.Type != events.UnknownReply {
			_ = handler.replyPublisher.Publish(reply)
		}

	}

}
