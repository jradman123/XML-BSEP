package dto

type UserActivateRequest struct {
	Username string `json:"username" validate:"required,min=2,max=30" `
	Code     string `json:"code" validate:"required,min=2,max=30"`
}
