package main

type CreatedUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
}
