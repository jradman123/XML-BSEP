package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"post/module/application"
	"post/module/infrastructure/api"
)

type UserCommandHandler struct {
	userService       *application.UserService
	postService       *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUserCommandHandler(userService *application.UserService, postService *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) (*UserCommandHandler, error) {
	o := &UserCommandHandler{
		userService:       userService,
		postService:       postService,
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
		err = handler.postService.UpdateUserPosts(user)
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

	case events.ActivateUser:
		err := handler.userService.ActivateUserAccount(user.UserId)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack)
		}
		reply = api.MapUserReply(user, events.UserActivated)
	case events.ChangeEmail:

	default:
		reply = api.MapUserReply(user, events.UnknownReply)
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
