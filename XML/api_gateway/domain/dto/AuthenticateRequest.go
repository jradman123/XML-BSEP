package dto

type AuthenticateRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Token    int    `json:"token" form:"token" binding:"required"`
}
