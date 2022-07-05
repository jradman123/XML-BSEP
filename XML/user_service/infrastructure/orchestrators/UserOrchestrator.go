package orchestrators

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/user_events"
	"fmt"
	"user/module/domain/model"
	"user/module/infrastructure/api"
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
	err = o.replySubscriber.Subscribe(o.handleConnection) //slusa odgovore connection servisa nadam se
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UserOrchestrator) CreateUser(user *model.User) error {

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

func (o *UserOrchestrator) handleConnection(reply *events.UserConnectionReply) events.UserReplyType {
	//TODO:We check what is the next command type
	//TODO:handle rollback if needed
	fmt.Println("BAAAAAAACKS connection")
	return events.UnknownReply

}

func (o *UserOrchestrator) CreateConnectionUser(user *model.User) error {

	events := &events.ConnectionUserCommand{
		Type: events.CreateUser,
		User: events.ConnectionUser{
			UserUID:       user.ID.String(),
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ProfileStatus: string(user.ProfileStatus),
			Interests:     api.MapToStringArrayInterests(user.Interests),
			Experiences:   api.MapToStringArrayExperiences(user.Experiences),
			Educations:    api.MapToStringArrayEducations(user.Educations),
			Skills:        api.MapToStringArraySkills(user.Skills),
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) EditConnectionUser(user *model.User) error {

	events := &events.ConnectionUserCommand{
		Type: events.UpdateUser,
		User: events.ConnectionUser{
			UserUID:       user.ID.String(),
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ProfileStatus: string(user.ProfileStatus),
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) EditConnectionUserProfessionalDetails(user *model.User) error {

	events := &events.ConnectionUserCommand{
		Type: events.UpdateUserProfessionalDetails,
		User: events.ConnectionUser{
			UserUID:       user.ID.String(),
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ProfileStatus: string(user.ProfileStatus),
			Interests:     api.MapToStringArrayInterests(user.Interests),
			Experiences:   api.MapToStringArrayExperiences(user.Experiences),
			Educations:    api.MapToStringArrayEducations(user.Educations),
			Skills:        api.MapToStringArraySkills(user.Skills),
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) DeleteConnectionUser(user *model.User) error {

	events := &events.ConnectionUserCommand{
		Type: events.DeleteUser,
		User: events.ConnectionUser{
			UserUID:       user.ID.String(),
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ProfileStatus: string(user.ProfileStatus),
			Interests:     api.MapToStringArrayInterests(user.Interests),
			Experiences:   api.MapToStringArrayExperiences(user.Experiences),
			Educations:    api.MapToStringArrayEducations(user.Educations),
			Skills:        api.MapToStringArraySkills(user.Skills),
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) ChangeProfileStatus(user *model.User) error {

	events := &events.ConnectionUserCommand{
		Type: events.ChangeProfileStatus,
		User: events.ConnectionUser{
			UserUID:       user.ID.String(),
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ProfileStatus: string(user.ProfileStatus),
			Interests:     api.MapToStringArrayInterests(user.Interests),
			Experiences:   api.MapToStringArrayExperiences(user.Experiences),
			Educations:    api.MapToStringArrayEducations(user.Educations),
			Skills:        api.MapToStringArraySkills(user.Skills),
		},
	}

	return o.commandPublisher.Publish(events)
}
