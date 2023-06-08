package dto

type CategoryCreateDto struct {
	Name string `json:"name" form:"category_name" binding:"required"`
}

type CategoryUpdateDto struct {
	Id   uint64 `json:"id" form:"category_id" binding:"required"`
	Name string `json:"name" form:"category_name" binding:"required"`
}
