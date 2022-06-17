package dto

type PasswordLessLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
}
