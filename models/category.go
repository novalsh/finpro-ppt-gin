package models

import "time"

type Category struct {
	CategoryId   uint64    `gorm:"primary_key:auto_increment" json:"category_id"`
	CategoryName string    `gorm:"type:varchar(255):not null" json:"category_name"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
