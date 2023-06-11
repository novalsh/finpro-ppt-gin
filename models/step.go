package models

import "time"

type Step struct {
	Id        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	TodoId    uint64    `gorm:"not null" json:"-"`
	Todo      Todo      `gorm:"foreignkey:TodoId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"todo"`
	Name      string    `gorm:"type:varchar(255):not null" json:"name"`
	Detail    string    `gorm:"type:varchar(255):not null" json:"detail"`
	Status    string    `gorm:"type:varchar(255):not null" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
