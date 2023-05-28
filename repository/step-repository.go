package repository

import (
	"gorm.io/gorm"

	"finpro/models"
)

type StepRepository interface {
	InsertStep(b models.Step) models.Step
	UpdateStep(b models.Step) models.Step
	DeleteStep(b models.Step)
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
	db.connection.Preload("TodoId").Find(&b)
	return b
}

func (db *stepConnection) DeleteStep(b models.Step) {
	db.connection.Delete(&b)
}

func (db *stepConnection) FindAllStep() []models.Step {
	var step []models.Step
	db.connection.Preload("TodoId").Find(&step)
	return step
}

func (db *stepConnection) FindStepById(stepId uint64) models.Step {
	var step models.Step
	db.connection.Preload("TodoId").Find(&step, stepId)
	return step
}
