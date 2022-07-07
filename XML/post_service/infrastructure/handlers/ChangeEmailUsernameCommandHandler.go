package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"fmt"
	"post/module/application"
	"post/module/infrastructure/api"
)

type ChangeEmailUsernameCommandHandler struct {
	userService       *application.UserService
	postService       *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewChangeEmailUsernameCommandHandler(userService *application.UserService, postService *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) (*ChangeEmailUsernameCommandHandler, error) {
	o := &ChangeEmailUsernameCommandHandler{
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

func (handler *ChangeEmailUsernameCommandHandler) handle(command *events.ChangeEmailUsernameCommand) {

	user := api.MapFromChangeEmailCommandToUser(command)
	var reply = &events.UserReply{}
	switch command.Type {
	case events.ChangeEmail:
		user, err := handler.userService.UpdateUser(user)
		if err != nil {
			reply = api.MapUserReply(user, events.ChangedEmailRolledBack)
		}
		reply = api.MapUserReply(user, events.ChangedEmail)
	case events.ChangeUsername:
		fmt.Println("usao u promjenu username")
		user, err := handler.userService.UpdateUser(user)
		if err != nil {
			reply = api.MapUserReply(user, events.ChangedUsernameRolledBack)
		}
		reply = api.MapUserReply(user, events.ChangedUsername)
	default:
		reply = api.MapUserReply(user, events.UnknownReply)
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
