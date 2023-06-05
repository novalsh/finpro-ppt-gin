package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(128);not null" json:"name"`
	Email     string    `gorm:"column:email;type:varchar(128);not null" json:"email"`
	Phone     string    `gorm:"column:phone;type:varchar(64);not null" json:"phone"`
	Password  string    `gorm:"column:password;type:varchar(256);not null" json:"password"`
	Token     string    `gorm:"column:token;type:varchar(256)" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
