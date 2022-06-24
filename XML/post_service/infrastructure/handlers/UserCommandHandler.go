package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"post/module/application"
	"post/module/infrastructure/api"
)

type UserCommandHandler struct {
	service           *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUserCommandHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*UserCommandHandler, error) {
	o := &UserCommandHandler{
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

func (handler *UserCommandHandler) handle(command *events.UserCommand) {

	user := api.MapNewUser(command)
	var reply = &events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		user, err := handler.service.CreateUser(user)
		if err != nil {
			reply = api.MapReplyUser(user, events.UserRolledBack)
		}
		reply = api.MapReplyUser(user, events.UserCreated)

	case events.UpdateUser:
		user, err := handler.service.UpdateUser(user)
		if err != nil {
			reply = api.MapReplyUser(user, events.UserRolledBack)
		}
		reply = api.MapReplyUser(user, events.UserUpdated)

	case events.DeleteUser:
		err := handler.service.DeleteUser(user.UserId)
		if err != nil {
			reply = api.MapReplyUser(user, events.UserRolledBack)
		}
		reply = api.MapReplyUser(user, events.UserDeleted)

	case events.ActivateUser:
		err := handler.service.ActivateUserAccount(user.UserId)
		if err != nil {
			reply = api.MapReplyUser(user, events.UserRolledBack)
		}
		reply = api.MapReplyUser(user, events.UserDeleted)
	default:
		reply = api.MapReplyUser(user, events.UnknownReply)
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
