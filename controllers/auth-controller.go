package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/helper"
	"finpro/models"
	"finpro/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) Logout(ctx *gin.Context) {
	// Extract the token from the request header or any other location
	// For example, if the token is passed in the "Authorization" header as a Bearer token:
	tokenString := ctx.GetHeader("Authorization")
	// Remove the "Bearer " prefix from the token string
	accessToken := strings.Replace(tokenString, "Bearer ", "", 1)

	// Perform any necessary token validation or additional logic

	// Assuming you have a method in your JWT service to validate and extract the user ID from the token
	_, err := c.jwtService.ValidateToken(accessToken)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// Return a success response
	response := helper.BuildResponse(true, "Logout successful", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
