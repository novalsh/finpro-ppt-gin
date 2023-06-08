package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/service"
)

type TodoController interface {
	InsertTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	FindAllTodo(ctx *gin.Context)
	FindTodoById(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(todoServ service.TodoService) TodoController {
	return &todoController{
		todoService: todoServ,
	}
}

func (controller *todoController) InsertTodo(ctx *gin.Context) {
	var todoCreateDto dto.TodoCreateDto
	err := ctx.ShouldBindJSON(&todoCreateDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk menyimpan todo
	createdTodo := controller.todoService.InsertTodo(todoCreateDto)

	ctx.JSON(http.StatusOK, createdTodo)
}

func (controller *todoController) UpdateTodo(ctx *gin.Context) {
	var todoUpdateDto dto.TodoUpdateDto
	err := ctx.ShouldBindJSON(&todoUpdateDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk mengupdate todo
	updatedTodo := controller.todoService.UpdateTodo(todoUpdateDto)

	ctx.JSON(http.StatusOK, updatedTodo)
}

func (controller *todoController) DeleteTodo(ctx *gin.Context) {
	// Ambil ID todo dari parameter URL
	TodoID := ctx.Param("id")

	// Konversi TodoID menjadi uint64
	todoIDUint, err := strconv.ParseUint(TodoID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// Panggil service untuk menghapus todo
	err = controller.todoService.DeleteTodoById(todoIDUint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func (controller *todoController) FindAllTodo(ctx *gin.Context) {
	// Panggil service untuk mendapatkan daftar semua todo
	todos := controller.todoService.FindAllTodo()

	ctx.JSON(http.StatusOK, todos)
}

func (controller *todoController) FindTodoById(ctx *gin.Context) {
	// Ambil ID todo dari parameter URL
	todoId := ctx.Param("id")

	TodoId, err := strconv.ParseUint(todoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk mendapatkan todo berdasarkan ID
	todo := controller.todoService.FindTodoById(TodoId)

	ctx.JSON(http.StatusOK, gin.H{"todo": todo})
}
