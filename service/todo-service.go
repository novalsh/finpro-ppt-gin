package service

import (
	"log"

	"github.com/mashingan/smapping"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type TodoService interface {
	InsertTodo(b dto.TodoCreateDto) models.Todo
	UpdateTodoById(b dto.TodoUpdateDto) models.Todo
	DeleteTodoById(b models.Todo)
	FindAllTodo() []models.Todo
	FindTodoById(todoID uint64) models.Todo
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) InsertTodo(b dto.TodoCreateDto) models.Todo {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.InsertTodo(todo)
	return res
}

func (service *todoService) UpdateTodoById(b dto.TodoUpdateDto) models.Todo {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.UpdateTodoById(todo)
	return res
}

func (service *todoService) DeleteTodoById(b models.Todo) {
	service.todoRepository.DeleteTodoById(b)
}

func (service *todoService) FindAllTodo() []models.Todo {
	return service.todoRepository.FindAllTodo() // Mengambil semua todo
}

func (service *todoService) FindTodoById(todoID uint64) models.Todo {
	return service.todoRepository.FindTodoById(todoID) // Mengambil todo berdasarkan ID
}
