package models

import (
	"time"
)

type Todo struct {
	TodoId              uint64    `gorm:"primary_key:auto_increment" json:"todo_id"`
	TodoName            string    `gorm:"type:varchar(255):not null" json:"todo_name"`
	TodoNote            string    `gorm:"type:longtext:not null" json:"todo_note"`
	UserId              User      `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	TodoDifficultyLevel uint16    `gorm:"type:int(11):not null" json:"todo_difficulty_level"`
	CategoryId          Category  `gorm:"foreignkey:CategoryId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category_id"`
	TodoLink            string    `gorm:"type:varchar(512):not null" json:"todo_link"`
	TodoDeadline        time.Time `gorm:"type:datetime:not null" json:"todo_deadline"`
	TodoWeight          uint16    `gorm:"type:int(11):not null" json:"todo_weight"`
	TodoDeadlineWeight  uint16    `gorm:"type:int(11):not null" json:"todo_deadline_weight"`
	TodoLevelWeight     uint16    `gorm:"type:int(11):not null" json:"todo_level_weight"`
	TodoCluster         uint16    `gorm:"type:int(11):not null" json:"todo_cluster_weight"`
	TodoStatus          string    `gorm:"type:varchar(255):not null" json:"todo_status"`
	CreatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
