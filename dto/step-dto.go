package dto

type StepCreateDto struct {
	TodoId     uint64 `json:"todo_id,omitempty" form:"todo_id,omitempty"`
	StepName   string `json:"step_name" form:"step_name" binding:"required"`
	StepDetail string `json:"step_detail" form:"step_detail" binding:"required"`
	StepStatus string `json:"step_status" form:"step_status" binding:"required"`
}

type StepUpdateDto struct {
	StepId     uint64 `json:"step_id" form:"step_id" binding:"required"`
	TodoId     uint64 `json:"todo_id,omitempty" form:"todo_id,omitempty"`
	StepName   string `json:"step_name" form:"step_name" binding:"required"`
	StepDetail string `json:"step_detail" form:"step_detail" binding:"required"`
	StepStatus string `json:"step_status" form:"step_status" binding:"required"`
}
