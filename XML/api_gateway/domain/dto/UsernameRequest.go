package dto

type UsernameRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
}
