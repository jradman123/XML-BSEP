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

type ConnectionUser struct {
	UserUID       string
	Username      string
	FirstName     string
	LastName      string
	ProfileStatus string
}

type ConnectionUserCommand struct {
	User ConnectionUser
	Type UserCommandType
}

type UserConnectionReply struct {
	Type UserReplyType
}
