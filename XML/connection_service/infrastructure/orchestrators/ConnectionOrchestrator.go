package orchestrators

import (
	events "common/module/saga/connection_events"
	saga "common/module/saga/messaging"
	"fmt"
)

type ConnectionOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewConnectionOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*ConnectionOrchestrator, error) {
	o := &ConnectionOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *ConnectionOrchestrator) Connect(userSender string, userReceiver string, status string) error {

	content := ""
	redirectTo := ""
	if status == "REQUEST_SENT" {
		content = userSender + " wants to connect with you."
		redirectTo = "/network#invitations"
	} else if status == "CONNECTED" {
		content = userSender + " connected with you."
		redirectTo = "/network#connections"
	}

	events := &events.ConnectionNotificationCommand{
		Type: events.Connect,
		Notification: events.Notification{
			Content:          content,
			RedirectPath:     redirectTo,
			NotificationFrom: userSender,
			NotificationTo:   userReceiver,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *ConnectionOrchestrator) AcceptConnection(userSender string, userReceiver string) error {

	events := &events.ConnectionNotificationCommand{
		Type: events.AcceptRequest,
		Notification: events.Notification{
			Content:          userSender + " accepted your invitation to connect.",
			RedirectPath:     "/public-profile/" + userSender,
			NotificationFrom: userSender,
			NotificationTo:   userReceiver,
		},
	}
	return o.commandPublisher.Publish(events)
}

func (o *ConnectionOrchestrator) handle(reply *events.ConnectionNotificationReply) events.ConnectionNotificationReplyType {
	if reply.Type == events.NotificationSent {
		fmt.Println("Senttttttttt")
	}
	if reply.Type == events.UnknownReply {
		fmt.Println("Unknown ")
	}
	fmt.Println("upsssss")
	return events.UnknownReply

}
