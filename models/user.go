package models

import "time"

type User struct {
	UserId        uint64    `gorm:"primary_key:auto_increment" json:"user_id"`
	UserName      string    `gorm:"type:varchar(255):not null" json:"user_name"`
	UserGmail     string    `gorm:"type:varchar(255):not null" json:"user_gmail"`
	UserGoogleId  string    `gorm:"type:varchar(255):not null" json:"user_google_id"`
	UserPicture   string    `gorm:"type:varchar(255):not null" json:"user_picture"`
	Password      string    `gorm:"type:varchar(255):not null" json:"password"`
	UserPronounce string    `gorm:"type:varchar(255):not null" json:"user_pronounce"`
	UserPhone     string    `gorm:"type:varchar(255):not null" json:"user_phone"`
	UserToken     string    `gorm:"type:varchar(255):not null" json:"user_token"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
