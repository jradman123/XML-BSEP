package handlers

import (
	events "common/module/saga/connection_events"
	saga "common/module/saga/messaging"
	"context"
	"message/module/application"
	"message/module/infrastructure/api"
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

func (handler *ConnectionNotificationCommandHandler) handle(command *events.ConnectionNotificationCommand) {

	notification := api.MapNewConnectionNotification(command)
	var reply = &events.ConnectionNotificationReply{}
	switch command.Type {
	case events.Connect:
		_, err := handler.service.Create(notification, context.TODO())
		if err != nil {
			reply = api.MapConnectionNotificationReply(events.NotificationNotSent)
		}
		reply = api.MapConnectionNotificationReply(events.NotificationSent)
	case events.AcceptRequest:
		_, err := handler.service.Create(notification, context.TODO())
		if err != nil {
			reply = api.MapConnectionNotificationReply(events.NotificationNotSent)
		}
		reply = api.MapConnectionNotificationReply(events.NotificationSent)
	default:
		reply = api.MapConnectionNotificationReply(events.UnknownReply)

		if reply.Type != events.UnknownReply {
			_ = handler.replyPublisher.Publish(reply)
		}

	}

}
