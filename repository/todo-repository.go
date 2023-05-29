package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type TodoRepository interface {
	InsertTodo(b models.Todo) models.Todo
	UpdateTodoById(b models.Todo) models.Todo
	DeleteTodoById(b models.Todo)
	FindAllTodo() []models.Todo
	FindTodoById(todoId uint64) models.Todo
	FindTodoByUserId(userId uint64) []models.Todo
	FindTodoByCategoryId(categoryId uint64) []models.Todo
}

type todoConnection struct {
	connection *gorm.DB
}

func NewTodoRepository(dbConn *gorm.DB) TodoRepository {
	return &todoConnection{
		connection: dbConn,
	}
}

func (db *todoConnection) InsertTodo(b models.Todo) models.Todo {
	db.connection.Save(&b)
	db.connection.Preload("UserId").Preload("CategoryId").Find(&b)
	return b
}

func (db *todoConnection) UpdateTodoById(b models.Todo) models.Todo {
	db.connection.Save(&b)
	db.connection.Preload("UserId").Preload("CategoryId").Find(&b)
	return b
}

func (db *todoConnection) DeleteTodoById(b models.Todo) {
	db.connection.Delete(&b)
}

func (db *todoConnection) FindAllTodo() []models.Todo {
	var todo []models.Todo
	db.connection.Preload("UserId").Preload("CategoryId").Find(&todo)
	return todo
}

func (db *todoConnection) FindTodoById(todoId uint64) models.Todo {
	var todo models.Todo
	db.connection.Preload("UserId").Preload("CategoryId").Find(&todo, todoId)
	return todo
}

func (db *todoConnection) FindTodoByUserId(userId uint64) []models.Todo {
	var todo []models.Todo
	db.connection.Preload("UserId").Preload("CategoryId").Find(&todo, "user_id = ?", userId)
	return todo
}

func (db *todoConnection) FindTodoByCategoryId(categoryId uint64) []models.Todo {
	var todo []models.Todo
	db.connection.Preload("UserId").Preload("CategoryId").Find(&todo, "category_id = ?", categoryId)
	return todo
}
