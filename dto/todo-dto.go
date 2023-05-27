package dto

type TodoCreateDto struct {
	TodoName            string `json:"todo_name" form:"todo_name" binding:"required"`
	TodoNote            string `json:"todo_note" form:"todo_note" binding:"required"`
	UserId              uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	TodoDifficultyLevel string `json:"todo_difficulty_level" form:"todo_difficulty_level" binding:"required"`
	CategoryId          uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
	TodoLink            string `json:"todo_link" form:"todo_link" binding:"required"`
	TodoDeadline        string `json:"todo_deadline" form:"todo_deadline" binding:"required"`
	TodoWeight          string `json:"todo_weight" form:"todo_weight" binding:"required"`
	TodoDeadlineWeight  string `json:"todo_deadline_weight" form:"todo_deadline_weight" binding:"required"`
	TodoLevelWeight     string `json:"todo_level_weight" form:"todo_level_weight" binding:"required"`
	TodoCluster         string `json:"todo_cluster" form:"todo_cluster" binding:"required"`
	TodoStatus          string `json:"todo_status" form:"todo_status" binding:"required"`
}

type TodoUpdateDto struct {
	TodoId              uint64 `json:"todo_id" form:"todo_id" binding:"required"`
	TodoName            string `json:"todo_name" form:"todo_name" binding:"required"`
	TodoNote            string `json:"todo_note" form:"todo_note" binding:"required"`
	UserId              uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	TodoDifficultyLevel string `json:"todo_difficulty_level" form:"todo_difficulty_level" binding:"required"`
	CategoryId          uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
	TodoLink            string `json:"todo_link" form:"todo_link" binding:"required"`
	TodoDeadline        string `json:"todo_deadline" form:"todo_deadline" binding:"required"`
	TodoWeight          string `json:"todo_weight" form:"todo_weight" binding:"required"`
	TodoDeadlineWeight  string `json:"todo_deadline_weight" form:"todo_deadline_weight" binding:"required"`
	TodoLevelWeight     string `json:"todo_level_weight" form:"todo_level_weight" binding:"required"`
	TodoCluster         string `json:"todo_cluster" form:"todo_cluster" binding:"required"`
	TodoStatus          string `json:"todo_status" form:"todo_status" binding:"required"`
}
