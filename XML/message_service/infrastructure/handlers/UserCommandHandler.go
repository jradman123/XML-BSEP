package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"message/module/application"
	"message/module/infrastructure/api"
)

type UserCommandHandler struct {
	userService       *application.UserService
	messageService    *application.MessageService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUserCommandHandler(userService *application.UserService, messageService *application.MessageService, publisher saga.Publisher, subscriber saga.Subscriber) (*UserCommandHandler, error) {
	o := &UserCommandHandler{
		userService:       userService,
		messageService:    messageService,
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
		user, err := handler.userService.CreateUser(user)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack)
		}
		reply = api.MapUserReply(user, events.UserCreated)

	case events.UpdateUser:
		user, err := handler.userService.UpdateUser(user)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack)
		}
		err = handler.messageService.UpdateUserMessages(user)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack)
		}
		reply = api.MapUserReply(user, events.UserUpdated)

	case events.DeleteUser:
		err := handler.userService.DeleteUser(user.UserId)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack)
		}
		reply = api.MapUserReply(user, events.UserDeleted)
	default:
		reply = api.MapUserReply(user, events.UnknownReply)
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
