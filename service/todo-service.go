package service

import (
	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type TodoService interface {
	InsertTodo(b dto.TodoCreateDto) dto.TodoCreateDto
	UpdateTodo(b dto.TodoUpdateDto) dto.TodoUpdateDto
	DeleteTodoById(todoID uint64) error
	FindAllTodo() []models.Todo
	FindTodoById(todoID uint64) models.Todo
	KMeans() []models.Todo
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) InsertTodo(b dto.TodoCreateDto) dto.TodoCreateDto {
	todoToInsert := models.Todo{}
	todoToInsert.UserId = b.UserId
	todoToInsert.CategoryId = b.CategoryId
	todoToInsert.Name = b.Name
	todoToInsert.Note = b.Note
	todoToInsert.Deadline = b.Deadline
	todoToInsert.Level = b.Level
	// todoToInsert.Cluster = b.Cluster
	todoInserted := service.todoRepository.InsertTodo(todoToInsert)

	// Konversi nilai todoInserted menjadi dto.TodoCreateDto
	insertedDto := dto.TodoCreateDto{
		UserId:     todoInserted.UserId,
		CategoryId: todoInserted.CategoryId,
		Name:       todoInserted.Name,
		Note:       todoInserted.Note,
		Deadline:   todoInserted.Deadline,
		Level:      todoInserted.Level,
		// Cluster:    todoInserted.Cluster,
	}

	return insertedDto
}

func (service *todoService) UpdateTodo(b dto.TodoUpdateDto) dto.TodoUpdateDto {
	todoToUpdate := models.Todo{}
	todoToUpdate.ID = b.Id
	todoToUpdate.UserId = b.UserId
	todoToUpdate.CategoryId = b.CategoryId
	todoToUpdate.Name = b.Name
	todoToUpdate.Note = b.Note
	todoToUpdate.Deadline = b.Deadline
	todoToUpdate.Level = b.Level
	// todoToUpdate.Cluster = b.Cluster
	todoUpdated := service.todoRepository.UpdateTodo(todoToUpdate)

	// Konversi nilai todoUpdated menjadi dto.TodoUpdateDto
	updatedDto := dto.TodoUpdateDto{
		Id:         todoUpdated.ID,
		UserId:     todoUpdated.UserId,
		CategoryId: todoUpdated.CategoryId,
		Name:       todoUpdated.Name,
		Note:       todoUpdated.Note,
		Deadline:   todoUpdated.Deadline,
		Level:      todoUpdated.Level,
		// Cluster:    todoUpdated.Cluster,
	}

	return updatedDto
}

func (service *todoService) DeleteTodoById(todoId uint64) error {
	err := service.todoRepository.DeleteTodoById(todoId)
	if err != nil {
		return err
	}
	return nil
}

func (service *todoService) FindAllTodo() []models.Todo {
	return service.todoRepository.FindAllTodo() // Mengambil semua todo
}

func (service *todoService) FindTodoById(todoID uint64) models.Todo {
	return service.todoRepository.FindTodoById(todoID) // Mengambil todo berdasarkan ID
}

func (service *todoService) KMeans() []models.Todo {
	return service.todoRepository.KMeans()
}
