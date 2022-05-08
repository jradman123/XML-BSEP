package dto

import "user/module/model"

type NewUser struct {
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	PhoneNumber string       `json:"phoneNumber"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Gender      model.Gender `json:"gender"`
}
