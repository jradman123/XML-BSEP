package post_events

type Notification struct {
	Content          string
	RedirectPath     string
	NotificationFrom string
	NotificationTo   string
}

type PostNotificationCommandType int8

const (
	LikePost PostNotificationCommandType = iota
	DislikePost
	CommentPost
	UnknownCommand
)

type PostNotificationReplyType int8

const (
	NotificationSent PostNotificationReplyType = iota
	NotificationNotSent
	UnknownReply
)

type PostNotificationCommand struct {
	Notification Notification
	Type         PostNotificationCommandType
}

type PostNotificationReply struct {
	Type PostNotificationReplyType
	// potential notyyyy
}
