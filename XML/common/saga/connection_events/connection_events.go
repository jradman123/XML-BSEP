package connection_events

type Notification struct {
	Content          string
	RedirectPath     string
	NotificationFrom string
	NotificationTo   string
}

type ConnectionNotificationCommandType int8

const (
	Connect ConnectionNotificationCommandType = iota
	AcceptRequest
)

type ConnectionNotificationReplyType int8

const (
	NotificationSent ConnectionNotificationReplyType = iota
	NotificationNotSent
	UnknownReply
)

type ConnectionNotificationCommand struct {
	Notification Notification
	Type         ConnectionNotificationCommandType
}

type ConnectionNotificationReply struct {
	Type ConnectionNotificationReplyType
}
