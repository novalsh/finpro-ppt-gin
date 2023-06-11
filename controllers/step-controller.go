package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/service"
)

type StepController interface {
	InsertStep(ctx *gin.Context)
	UpdateStep(ctx *gin.Context)
	DeleteStep(ctx *gin.Context)
	FindAllStep(ctx *gin.Context)
	FindStepById(ctx *gin.Context)
}

type stepController struct {
	stepService service.StepService
}

func NewStepController(stepServ service.StepService) StepController {
	return &stepController{
		stepService: stepServ,
	}
}

func (controller *stepController) InsertStep(ctx *gin.Context) {
	var StepCreateDtp dto.StepCreateDto
	err := ctx.ShouldBindJSON(&StepCreateDtp)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk menyimpan Step
	createdStep := controller.stepService.InsertStep(StepCreateDtp)

	ctx.JSON(200, createdStep)
}

func (controller *stepController) UpdateStep(ctx *gin.Context) {
	var StepUpdateDto dto.StepUpdateDto
	err := ctx.ShouldBindJSON(&StepUpdateDto)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk mengupdate Step
	updatedStep := controller.stepService.UpdateStep(StepUpdateDto)

	ctx.JSON(200, updatedStep)
}

func (controller *stepController) DeleteStep(ctx *gin.Context) {
	StepID := ctx.Param("id")

	StepIDUint, err := strconv.ParseUint(StepID, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid step id"})
		return
	}

	err = controller.stepService.DeleteStepById(StepIDUint)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

func (controller *stepController) FindAllStep(ctx *gin.Context) {
	Steps := controller.stepService.FindAllStep()

	ctx.JSON(200, Steps)
}

func (controller *stepController) FindStepById(ctx *gin.Context) {
	StepID := ctx.Param("id")

	StepIDUint, err := strconv.ParseUint(StepID, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid step id"})
		return
	}

	Step := controller.stepService.FindStepById(StepIDUint)

	ctx.JSON(200, Step)
}
