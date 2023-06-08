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

type CategoryController interface {
	AllCategory(ctx *gin.Context)
	FindCategoryById(ctx *gin.Context)
	InsertCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	GetCategoryByToken(token string) string
}

type categoryController struct {
	categoryService service.CategoryService
	jwtService      service.JWTService
}

func NewCategoryController(categoryServ service.CategoryService, jwtService service.JWTService) CategoryController {
	return &categoryController{
		categoryService: categoryServ,
		jwtService:      jwtService,
	}
}

func (c *categoryController) AllCategory(ctx *gin.Context) {
	var category []models.Category = c.categoryService.FindAllCategory()
	res := helper.BuildResponse(true, "OK!", category)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) FindCategoryById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var category models.Category = c.categoryService.FindCategoryById(id)
	if (category == models.Category{}) {
		res := helper.BuildErrorResponse("Failed to get category", "Category not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", category)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *categoryController) InsertCategory(ctx *gin.Context) {
	var categoryCreateDto dto.CategoryCreateDto
	errDTO := ctx.ShouldBind(&categoryCreateDto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.categoryService.InsertCategory(categoryCreateDto)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	var categoryUpdateDto dto.CategoryUpdateDto

	errDTO := ctx.ShouldBind(&categoryUpdateDto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.categoryService.UpdateCategory(categoryUpdateDto)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	var category models.Category
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	category.Id = id
	c.categoryService.DeleteCategory(category)
	res := helper.BuildResponse(true, "Deleted!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) GetCategoryByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
