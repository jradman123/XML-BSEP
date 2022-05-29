package dto

type NewRecoveryPasswordRequest struct {
	Username    string `json:"username" validate:"required,min=2,max=30" `
	NewPassword string `json:"password" validate:"required,min=10,max=30"`
	Code        string `json:"code" validate:"required,min=2,max=30"`
}
