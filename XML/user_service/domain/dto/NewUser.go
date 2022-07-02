package dto

//register new user dto
type NewUser struct {
	Username      string `json:"username"`
	Password      string `json:"password" validate:"required,min=10,max=30"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phoneNumber" validate:"required"`
	FirstName     string `json:"firstName" validate:"required,alpha,min=2,max=20"`
	LastName      string `json:"lastName" validate:"required,alpha,min=2,max=35"`
	Gender        string `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	DateOfBirth   string `json:"dateOfBirth" validate:"required"`
	RecoveryEmail string `json:"recoveryEmail" validate:"required,email"`
	ProfileStatus string `json:"profileStatus"`
}
