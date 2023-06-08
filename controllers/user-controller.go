package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/helper"
	"finpro/service"
)

type UserController interface {
	UpdateUser(ctx *gin.Context)
	ProfileUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userServ service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtService,
	}
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	result := c.userService.UpdateUser(userUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) ProfileUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	result := c.userService.ProfileUser(userID)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}
