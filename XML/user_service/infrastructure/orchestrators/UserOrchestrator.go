package orchestrators

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"fmt"
	"user/module/domain/model"
)

type UserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UserOrchestrator, error) {
	o := &UserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle) //slusa odgovore
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UserOrchestrator) CreateUser(user *model.User) error {

	fmt.Println("evo me u orchestratoru u create user ")

	events := &events.UserCommand{
		Type: events.CreateUser,
		User: events.User{
			UserId:    user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) ActivateUserAccount(user *model.User) error {
	events := events.UserCommand{
		Type: events.ActivateUser,
		User: events.User{
			UserId:    user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	return o.commandPublisher.Publish(events)
}
func (o *UserOrchestrator) UpdateUser(user *model.User) error {
	events := events.UserCommand{
		Type: events.UpdateUser,
		User: events.User{
			UserId:    user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) handle(reply *events.UserReply) events.UserReplyType {
	if reply.Type == events.UserRolledBack {
		fmt.Println("UserRolledBack")
	}
	if reply.Type == events.UserCreated {
		fmt.Println("UserCreated")
	}
	if reply.Type == events.UserUpdated {
		fmt.Println("UserUpdated")
	}
	fmt.Println("BAAAAAAACKS")
	return events.UnknownReply

}
