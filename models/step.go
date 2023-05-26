package models

import "time"

type Step struct {
	StepId     uint64    `gorm:"primary_key:auto_increment" json:"step_id"`
	TodoId     Todo      `gorm:"foreignkey:TodoId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"todo_id"`
	StepName   string    `gorm:"type:varchar(255):not null" json:"step_name"`
	StepDetail string    `gorm:"type:varchar(255):not null" json:"step_detail"`
	StepStatus string    `gorm:"type:varchar(255):not null" json:"step_status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
