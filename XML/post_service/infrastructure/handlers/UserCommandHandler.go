package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"context"
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
	user := api.MapNewUser(command, context.TODO())
	var reply = &events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		user, err := handler.userService.CreateUser(user, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		reply = api.MapUserReply(user, events.UserCreated, context.TODO())

	case events.UpdateUser:
		_, err := handler.userService.UpdateUser(user, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		err = handler.postService.UpdateUserPosts(user, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		reply = api.MapUserReply(user, events.UserUpdated, context.TODO())

	case events.DeleteUser:
		err := handler.userService.DeleteUser(user.UserId, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		reply = api.MapUserReply(user, events.UserDeleted, context.TODO())

	case events.ActivateUser:
		err := handler.userService.ActivateUserAccount(user.UserId, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		reply = api.MapUserReply(user, events.UserActivated, context.TODO())
	case events.ChangeEmail:
		_, err := handler.userService.UpdateUser(user, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.ChangedEmailRolledBack, context.TODO())
		}
		reply = api.MapUserReply(user, events.ChangedEmail, context.TODO())

	default:
		reply = api.MapUserReply(user, events.UnknownReply, context.TODO())
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
