package dto

//register new user dto
type NewRecoveryPasword struct {
	Username    string `json:"username" validate:"required,min=2,max=30" `
	NewPassword string `json:"newPassword" validate:"required,min=10,max=30"`
	Code        string `json:"code" validate:"required,min=10,max=30"`
}
