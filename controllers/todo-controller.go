package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/helper"
	"finpro/models"
	"finpro/service"
)

type TodoController interface {
	InsertTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	AllTodo(ctx *gin.Context)
	FindTodoById(ctx *gin.Context)
	GetUserIdByToken(token string) string
}

type todoController struct {
	todoService service.TodoService
	jwtService  service.JWTService
}

func NewTodoController(todoServ service.TodoService, jwtService service.JWTService) TodoController {
	return &todoController{
		todoService: todoServ,
		jwtService:  jwtService,
	}
}

func (c *todoController) InsertTodo(ctx *gin.Context) {
	var todoCreateDto dto.TodoCreateDto
	errDTO := ctx.ShouldBind(&todoCreateDto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userId := c.GetUserIdByToken(authHeader)
		convertedUserId, err := strconv.ParseUint(userId, 10, 64)
		if err == nil {
			todoCreateDto.UserId = convertedUserId
		}
		result := c.todoService.InsertTodo(todoCreateDto)
		response := helper.BuildResponse(true, "OK!", result)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *todoController) UpdateTodo(ctx *gin.Context) {
	var TodoUpdateDto dto.TodoUpdateDto
	errDTO := ctx.ShouldBind(&TodoUpdateDto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	if c.todoService.IsAllowedToEdit(id, TodoUpdateDto.Id) {
		id, errID := strconv.ParseUint(id, 10, 64)
		if errID == nil {
			TodoUpdateDto.UserId = id
		}
		result := c.todoService.UpdateTodo(TodoUpdateDto)
		response := helper.BuildResponse(true, "OK!", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You don't have permission to edit this todo", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) DeleteTodo(ctx *gin.Context) {
	var todo models.Todo
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	todo.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	if c.todoService.IsAllowedToEdit(userId, todo.ID) {
		c.todoService.DeleteTodo(todo)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You don't have permission to delete this todo", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) AllTodo(ctx *gin.Context) {
	var todos []models.Todo = c.todoService.FindAllTodo()
	res := helper.BuildResponse(true, "OK!", todos)
	ctx.JSON(http.StatusOK, res)
}

func (c *todoController) FindTodoById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var todo models.Todo = c.todoService.FindTodoById(id)
	if (todo == models.Todo{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", todo)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *todoController) GetUserIdByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id

}
