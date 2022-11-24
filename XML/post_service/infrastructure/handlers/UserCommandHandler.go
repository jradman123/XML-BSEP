package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"context"
	tracer "monitoring/module"
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

func (handler *UserCommandHandler) handle(command *events.UserCommand, ctx context.Context) {
	span := tracer.StartSpanFromContextMetadata(ctx, "userCommandHandler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	user := api.MapNewUser(command, ctx)
	var reply = &events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		user, err := handler.userService.CreateUser(user, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		reply = api.MapUserReply(user, events.UserCreated, ctx)

	case events.UpdateUser:
		_, err := handler.userService.UpdateUser(user, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		err = handler.postService.UpdateUserPosts(user, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		reply = api.MapUserReply(user, events.UserUpdated, ctx)

	case events.DeleteUser:
		err := handler.userService.DeleteUser(user.UserId, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		reply = api.MapUserReply(user, events.UserDeleted, ctx)

	case events.ActivateUser:
		err := handler.userService.ActivateUserAccount(user.UserId, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		reply = api.MapUserReply(user, events.UserActivated, ctx)
	case events.ChangeEmail:
		_, err := handler.userService.UpdateUser(user, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.ChangedEmailRolledBack, ctx)
		}
		reply = api.MapUserReply(user, events.ChangedEmail, ctx)

	default:
		reply = api.MapUserReply(user, events.UnknownReply, ctx)
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
