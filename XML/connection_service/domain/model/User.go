package model

type User struct {
	UserUID   string
	Username  string
	FirstName string
	LastName  string
	Status    ProfileStatus
}

type UserDetails struct {
	Interests   []string
	Skills      []string
	Educations  []string
	Experiences []string
}

type ProfileStatus string

const (
	Private ProfileStatus = "PRIVATE"
	Public  ProfileStatus = "PUBLIC"
)
