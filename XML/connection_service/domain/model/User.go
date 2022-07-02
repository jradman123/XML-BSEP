package model

type User struct {
	UserUID   string
	Username  string
	FirstName string
	LastName  string
	Status    ProfileStatus
}

type ProfileStatus string

const (
	Private ProfileStatus = "PRIVATE"
	Public  ProfileStatus = "PUBLIC"
)
