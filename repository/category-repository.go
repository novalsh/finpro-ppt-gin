package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type CategoryRepository interface {
	InsertCategory(categoryName string) models.Category
	UpdateCategory(categoryId uint64, categoryName string) models.Category
	DeleteCategory(categoryId uint64) models.Category
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

func (db *categoryConnection) InsertCategory(categoryName string) models.Category {
	category := models.Category{CategoryName: categoryName}
	db.connection.Save(&category)
	return category
}

func (db *categoryConnection) UpdateCategory(categoryId uint64, categoryName string) models.Category {
	var category models.Category
	db.connection.Find(&category, categoryId)
	category.CategoryName = categoryName
	db.connection.Save(&category)
	return category
}

func (db *categoryConnection) DeleteCategory(categoryId uint64) models.Category {
	var category models.Category
	db.connection.Find(&category, categoryId)
	db.connection.Delete(&category)
	return category
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
