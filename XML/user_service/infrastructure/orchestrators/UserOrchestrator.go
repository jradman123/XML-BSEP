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

	gender := ""
	if user.Gender == model.FEMALE {
		gender = "FEMALE"
	} else {
		gender = "MALE"
	}

	events := &events.UserCommand{
		Type: events.CreateUser,
		User: events.User{
			Username:      user.Username,
			Password:      user.Password,
			Email:         user.Email,
			PhoneNumber:   user.PhoneNumber,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Gender:        gender,
			DateOfBirth:   user.DateOfBirth.Format("01-02-2006"),
			RecoveryEmail: user.RecoveryEmail,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) handle(reply *events.UserReply) events.UserReplyType {
	//TODO:We check what is the next command type
	fmt.Println("BAAAAAAACKS")
	return events.UnknownReply

}
