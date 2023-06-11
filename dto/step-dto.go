package dto

type StepCreateDto struct {
	TodoId uint64 `json:"todo_id,omitempty" form:"todo_id,omitempty"`
	Name   string `json:"name" form:"name" binding:"required"`
	Detail string `json:"step_detail" form:"step_detail" binding:"required"`
	Status string `json:"step_status" form:"step_status" binding:"required"`
}

type StepUpdateDto struct {
	StepId uint64 `json:"id" form:"step_id" binding:"required"`
	TodoId uint64 `json:"todo_id,omitempty" form:"todo_id,omitempty"`
	Name   string `json:"name" form:"name" binding:"required"`
	Detail string `json:"step_detail" form:"step_detail" binding:"required"`
	Status string `json:"step_status" form:"step_status" binding:"required"`
}
