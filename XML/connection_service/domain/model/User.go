package model

type User struct {
	UserUID string
	Status  ProfileStatus
}

type ProfileStatus string

const (
	Private ProfileStatus = "PRIVATE"
	Public                = "PUBLIC"
)
