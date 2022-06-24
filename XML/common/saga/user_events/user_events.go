package user_events

type User struct {
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	FirstName     string
	LastName      string
	Gender        string
	DateOfBirth   string
	RecoveryEmail string
}

type UserCommandType int8

const (
	UpdateUser UserCommandType = iota
	CreateUser
	DeleteUser
	RollbackUser
	UnknownCommand
)

type UserCommand struct {
	User User
	Type UserCommandType
}

type UserReplyType int8

const (
	UserUpdated UserReplyType = iota
	UserRolledBack
	UserDeleted
	UserCreated
	UnknownReply
)

type UserReply struct {
	Type UserReplyType
}
