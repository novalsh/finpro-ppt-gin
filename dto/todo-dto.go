package dto

type TodoCreateDto struct {
	UserId     uint64 `json:"user_id" form:"user_id" binding:"required"`
	CategoryId uint64 `json:"category_id" form:"category_id" binding:"required"`
	Name       string `json:"name" form:"todo_name" binding:"required"`
	Note       string `json:"note" form:"todo_note" binding:"required"`
	Deadline   string `json:"deadline" form:"todo_deadline" binding:"required"`
	Level      string `json:"level" form:"todo_difficulty_level" binding:"required"`
	Cluster    string `json:"cluster" form:"todo_cluster" binding:"required"`
}

type TodoUpdateDto struct {
	Id         uint64 `json:"id" form:"id" binding:"required"`
	UserId     uint64 `json:"user_id" form:"user_id" binding:"required"`
	CategoryId uint64 `json:"category_id" form:"category_id" binding:"required"`
	Name       string `json:"name" form:"todo_name" binding:"required"`
	Note       string `json:"note" form:"todo_note" binding:"required"`
	Deadline   string `json:"deadline" form:"todo_deadline" binding:"required"`
	Level      string `json:"level" form:"todo_difficulty_level" binding:"required"`
	Cluster    string `json:"cluster" form:"todo_cluster" binding:"required"`
}
