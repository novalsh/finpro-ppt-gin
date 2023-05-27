package dto

type CategoryCreateDto struct {
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
}

type CategoryUpdateDto struct {
	CategoryId   uint64 `json:"category_id" form:"category_id" binding:"required"`
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
}
