package models

import "time"

type Todo struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserId     uint64    `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	CategoryId uint64    `gorm:"foreignkey:CategoryId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category_id"`
	Name       string    `gorm:"type:varchar(255):not null" json:"name"`
	Note       string    `gorm:"type:varchar(255):not null" json:"note"`
	Deadline   string    `gorm:"type:varchar(255):not null" json:"deadline"`
	Level      string    `gorm:"type:varchar(255):not null" json:"level"`
	Cluster    string    `gorm:"type:varchar(255):not null" json:"cluster"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
