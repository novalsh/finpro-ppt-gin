package dto

type UserUpdateDTO struct {
	UserId    uint64 `json:"user_id" form:"user_id" binding:"required"`
	UserName  string `json:"user_name" form:"user_name" binding:"required"`
	UserGmail string `json:"user_gmail" form:"user_gmail" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}
