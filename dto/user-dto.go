package dto

type UserUpdateDTO struct {
	Id       uint64 `json:"id" form:"user_id" binding:"required"`
	Name     string `json:"name" form:"user_name" binding:"required"`
	Email    string `json:"email" form:"user_gmail" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
