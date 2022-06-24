package dto

type AuthenticateResponse struct {
	Username string `json:"username" form:"username" binding:"required"`
	TwoFa    bool   `json:"twofa" form:"twofa" binding:"required"`
}
