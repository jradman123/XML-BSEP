package dto

type UserPersonalDetails struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	FirstName   string `json:"firstName" validate:"required,alpha,min=2,max=20"`
	LastName    string `json:"lastName" validate:"required,alpha,min=2,max=35"`
	Gender      string `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
	Biography   string `json:"biography"`
}
