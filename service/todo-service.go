package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type TodoService interface {
	InsertTodo(b dto.TodoCreateDto) models.Todo
	UpdateTodo(b dto.TodoUpdateDto) models.Todo
	DeleteTodo(b models.Todo)
	FindAllTodo() []models.Todo
	FindTodoById(todoID uint64) models.Todo
	IsAllowedToEdit(UserId string, ID uint64) bool
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

func (service *todoService) UpdateTodo(b dto.TodoUpdateDto) models.Todo {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.UpdateTodoById(todo)
	return res
}

func (service *todoService) DeleteTodo(b models.Todo) {
	service.todoRepository.DeleteTodoById(b)
}

func (service *todoService) FindAllTodo() []models.Todo {
	return service.todoRepository.FindAllTodo() // Mengambil semua todo
}

func (service *todoService) FindTodoById(todoID uint64) models.Todo {
	return service.todoRepository.FindTodoById(todoID) // Mengambil todo berdasarkan ID
}

func (service *todoService) IsAllowedToEdit(userID string, todoID uint64) bool {
	b := service.todoRepository.FindTodoById(todoID)
	id := fmt.Sprintf("%v", b.UserId)
	return userID == id
}
