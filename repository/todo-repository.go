package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type TodoRepository interface {
	InsertTodo(b models.Todo) models.Todo
	UpdateTodo(b models.Todo) models.Todo
	DeleteTodoById(id uint64) error
	FindAllTodo() []models.Todo
	FindTodoById(todoId uint64) models.Todo
	KMeans() []models.Todo
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
	return b
}

func (db *todoConnection) UpdateTodo(b models.Todo) models.Todo {
	db.connection.Save(&b)
	return b
}

func (db *todoConnection) DeleteTodoById(id uint64) error {
	todo := models.Todo{}
	result := db.connection.Delete(&todo, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *todoConnection) FindAllTodo() []models.Todo {
	var todo []models.Todo
	db.connection.Preload("User").Preload("Category").Find(&todo)
	return todo
}

func (db *todoConnection) FindTodoById(todoId uint64) models.Todo {
	var todo models.Todo
	db.connection.Preload("User").Preload("Category").Find(&todo, todoId)
	return todo
}

func (db *todoConnection) KMeans() []models.Todo {
	var todo []models.Todo
	db.connection.Preload("User").Find(&todo)
	return todo
}
