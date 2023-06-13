package models

import "time"

type Todo struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserId     uint64    `gorm:"not null" json:"-"`
	User       User      `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CategoryId uint64    `gorm:"not null" json:"-"`
	Category   Category  `gorm:"foreignkey:CategoryId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category_id"`
	Name       string    `gorm:"type:varchar(255):not null" json:"name"`
	Note       string    `gorm:"type:varchar(255)" json:"note"`
	Deadline   string    `gorm:"type:varchar(255):not null" json:"deadline"`
	Level      float64   `gorm:"type:varchar(255):not null" json:"level"`
	Cluster    string    `gorm:"type:varchar(255)" json:"cluster"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Todo) TableName() string {
	return "todos"
}
