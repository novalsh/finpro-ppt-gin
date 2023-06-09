package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type StepRepository interface {
	InsertStep(b models.Step) models.Step
	UpdateStep(b models.Step) models.Step
	DeleteStepById(stepID uint64) error
	FindAllStep() []models.Step
	FindStepById(stepId uint64) models.Step
}

type stepConnection struct {
	connection *gorm.DB
}

func NewStepRepository(dbConn *gorm.DB) StepRepository {
	return &stepConnection{
		connection: dbConn,
	}
}

func (db *stepConnection) InsertStep(b models.Step) models.Step {
	db.connection.Save(&b)
	db.connection.Preload("TodoId").Find(&b)
	return b

}

func (db *stepConnection) UpdateStep(b models.Step) models.Step {
	db.connection.Save(&b)
	return b
}

func (db *stepConnection) DeleteStepById(stepId uint64) error {
	step := models.Step{}
	result := db.connection.Delete(&step, stepId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *stepConnection) FindAllStep() []models.Step {
	var steps []models.Step
	db.connection.Preload("Todo").Find(&steps)
	return steps
}

func (db *stepConnection) FindStepById(stepId uint64) models.Step {
	var step models.Step
	db.connection.Preload("Todo").Find(&step, stepId)
	return step
}
