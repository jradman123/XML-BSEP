package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"context"
	"fmt"
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
	fmt.Println("evo me u command handleru u mess servisu prije caseova")

	user := api.MapNewUser(command, context.TODO())
	fmt.Sprintln(user)
	var reply = &events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		fmt.Println("evo me u command handleru u mess servisu case CREATE USER")
		user, err := handler.userService.CreateUser(user, context.TODO())
		if err != nil {
			fmt.Println("rolbekujerm jer error ")
			fmt.Println(err)
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		fmt.Println("nije rolbekovo")
		reply = api.MapUserReply(user, events.UserCreated, context.TODO())

	case events.UpdateUser:
		_, err := handler.userService.UpdateUser(user, context.TODO())
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, context.TODO())
		}
		err = handler.messageService.UpdateUserMessages(user, context.TODO())
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
	fmt.Println("evo me na kraj command handlera u mess servisu")

}
