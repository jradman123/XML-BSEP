package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/application"
	"post/module/domain/model"
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

	user := model.User{
		Id:        primitive.NewObjectID(),
		Username:  command.User.Username,
		FirstName: command.User.FirstName,
		LastName:  command.User.LastName,
	}
	var reply = events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		err := handler.service.CreateUser(user)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserCreated

		//TODO:Cannot update users' username
	case events.UpdateUser:
		err := handler.service.UpdateUser(user)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserUpdated

	case events.DeleteUser:
		err := handler.service.DeleteUser(user.Username)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserDeleted

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
