package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"connection/module/application/services"
	"connection/module/domain/model"
	"context"
	"fmt"
	tracer "monitoring/module"
)

type UserCommandHandler struct {
	service           *services.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUserCommandHandler(service *services.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*UserCommandHandler, error) {
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

func (handler *UserCommandHandler) handle(command *events.ConnectionUserCommand, ctx context.Context) {
	span := tracer.StartSpanFromContextMetadata(ctx, "userCommandHandler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	fmt.Println("usao u user command handler connection servisa")
	status := model.Public
	if command.User.ProfileStatus == "PRIVATE" {
		status = model.Private
	}

	user := model.User{
		UserUID:   command.User.UserUID,
		Username:  command.User.Username,
		FirstName: command.User.FirstName,
		LastName:  command.User.LastName,
		Status:    status,
	}

	var reply = events.UserConnectionReply{}
	switch command.Type {
	case events.CreateUser:
		err := handler.service.CreateUser(user, ctx)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserCreated

		// TODO:Cannot update users' username
	case events.UpdateUser:
		fmt.Println("events.UpdateUser")
		err := handler.service.UpdateUser(user, ctx)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserUpdated

	case events.DeleteUser:
		err := handler.service.DeleteUser(user, ctx)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.UserDeleted
	case events.ChangeProfileStatus:
		err := handler.service.ChangeProfileStatus(user, ctx)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.ProfileStatusChanged
	case events.UpdateUserProfessionalDetails:
		details := model.UserDetails{
			Skills:      command.User.Skills,
			Educations:  command.User.Educations,
			Interests:   command.User.Interests,
			Experiences: command.User.Experiences,
		}
		fmt.Println("evo me u handleru u connection servisu")
		fmt.Println(user)
		fmt.Println(details)
		err := handler.service.UpdateUserProfessionalDetails(user, details, ctx)
		if err != nil {
			reply.Type = events.UserRolledBack
		}
		reply.Type = events.ProfileStatusChanged
		fmt.Println("events.UpdateUserProfessionalDetails")
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}

	fmt.Println("dosao do kraja user command handler-a connection servisa")

}
