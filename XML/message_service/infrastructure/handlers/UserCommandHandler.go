package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"context"
	"fmt"
	"message/module/application"
	"message/module/infrastructure/api"
	tracer "monitoring/module"
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

func (handler *UserCommandHandler) handle(command *events.UserCommand, ctx context.Context) {
	span := tracer.StartSpanFromContextMetadata(ctx, "userCommandHandler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	fmt.Println("evo me u command handleru u mess servisu prije caseova")

	user := api.MapNewUser(command, ctx)
	fmt.Sprintln(user)
	var reply = &events.UserReply{}
	switch command.Type {
	case events.CreateUser:
		fmt.Println("evo me u command handleru u mess servisu case CREATE USER")
		user, err := handler.userService.CreateUser(user, ctx)
		if err != nil {
			fmt.Println("rolbekujerm jer error ")
			fmt.Println(err)
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		fmt.Println("nije rolbekovo")
		reply = api.MapUserReply(user, events.UserCreated, ctx)

	case events.UpdateUser:
		_, err := handler.userService.UpdateUser(user, ctx)
		if err != nil {
			reply = api.MapUserReply(user, events.UserRolledBack, ctx)
		}
		err = handler.messageService.UpdateUserMessages(user, ctx)
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
	fmt.Println("evo me na kraj command handlera u mess servisu")

}
