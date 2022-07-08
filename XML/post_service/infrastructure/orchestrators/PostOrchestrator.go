package orchestrators

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/post_events"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewPostOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*PostOrchestrator, error) {
	o := &PostOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *PostOrchestrator) LikePost(postId primitive.ObjectID, liker string, postOwner string) error {

	events := &events.PostNotificationCommand{
		Type: events.LikePost,
		Notification: events.Notification{
			Content:          liker + " liked your post.",
			RedirectPath:     "/post/" + postId.Hex(),
			NotificationFrom: liker,
			NotificationTo:   postOwner,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *PostOrchestrator) DislikePost(postId primitive.ObjectID, hater string, postOwner string) error {

	events := &events.PostNotificationCommand{
		Type: events.DislikePost,
		Notification: events.Notification{
			Content:          hater + " disliked your post.",
			RedirectPath:     "/post/" + postId.Hex(),
			NotificationFrom: hater,
			NotificationTo:   postOwner,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *PostOrchestrator) CommentPost(postId primitive.ObjectID, commenter string, postOwner string) error {

	events := &events.PostNotificationCommand{
		Type: events.CommentPost,
		Notification: events.Notification{
			Content:          commenter + " left a comment on your post.",
			RedirectPath:     "/post/" + postId.Hex(),
			NotificationFrom: commenter,
			NotificationTo:   postOwner,
		},
	}

	return o.commandPublisher.Publish(events)
}

func (o *PostOrchestrator) handle(reply *events.PostNotificationReply) events.PostNotificationReplyType {
	if reply.Type == events.NotificationSent {
		fmt.Println("Senttttttttt")
	}
	if reply.Type == events.UnknownReply {
		fmt.Println("Unknown ")
	}
	fmt.Println("upsssss")
	return events.UnknownReply

}
