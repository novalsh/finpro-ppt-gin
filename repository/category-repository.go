package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type CategoryRepository interface {
	InsertCategory(b models.Category) models.Category
	UpdateCategory(b models.Category) models.Category
	DeleteCategory(b models.Category)
	FindAllCategory() []models.Category
	FindCategoryById(categoryId uint64) models.Category
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbConn,
	}
}

func (db *categoryConnection) InsertCategory(b models.Category) models.Category {
	db.connection.Save(&b)
	return b
}

func (db *categoryConnection) UpdateCategory(b models.Category) models.Category {
	db.connection.Save(&b)
	return b
}

func (db *categoryConnection) DeleteCategory(b models.Category) {
	db.connection.Delete(b)
}

func (db *categoryConnection) FindAllCategory() []models.Category {
	var category []models.Category
	db.connection.Find(&category)
	return category
}

func (db *categoryConnection) FindCategoryById(categoryId uint64) models.Category {
	var category models.Category
	db.connection.Find(&category, categoryId)
	return category
}
