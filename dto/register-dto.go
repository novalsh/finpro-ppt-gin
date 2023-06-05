package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
