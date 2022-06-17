package dto

type Enable2FaResponse struct {
	Res bool   `json:"res" form:"res" binding:"required"`
	Uri string `json:"uri" form:"uri" binding:"required"`
}
