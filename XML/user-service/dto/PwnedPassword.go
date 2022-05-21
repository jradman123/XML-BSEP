package dto

type PwnedPasswordRequest struct {
	PwnedPassword string `json:"pwnedPassword"  validate:"required,min=10,max=30"`
}

type PwnedResponse struct {
	IsPwned bool   `json:"isPwned" `
	Message string `json:"message" `
}
