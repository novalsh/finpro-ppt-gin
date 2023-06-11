package service

import (
	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type StepService interface {
	InsertStep(b dto.StepCreateDto) dto.StepCreateDto
	UpdateStep(b dto.StepUpdateDto) dto.StepUpdateDto
	DeleteStepById(stepID uint64) error
	FindAllStep() []models.Step
	FindStepById(stepID uint64) models.Step
}

type stepService struct {
	stepRepository repository.StepRepository
}

func NewStepService(stepRepo repository.StepRepository) StepService {
	return &stepService{
		stepRepository: stepRepo,
	}
}

func (service *stepService) InsertStep(b dto.StepCreateDto) dto.StepCreateDto {
	StepToInsert := models.Step{}
	StepToInsert.TodoId = uint64(b.TodoId)
	StepToInsert.Name = b.Name
	StepToInsert.Detail = b.Detail
	StepToInsert.Status = b.Status
	StepInserted := service.stepRepository.InsertStep(StepToInsert)

	insertDto := dto.StepCreateDto{
		TodoId: uint64(StepInserted.TodoId),
		Name:   StepInserted.Name,
		Detail: StepInserted.Detail,
		Status: StepInserted.Status,
	}
	return insertDto
}

func (service *stepService) UpdateStep(b dto.StepUpdateDto) dto.StepUpdateDto {
	stepToUpdate := models.Step{}
	stepToUpdate.Id = b.StepId
	stepToUpdate.TodoId = uint64(b.TodoId)
	stepToUpdate.Name = b.Name
	stepToUpdate.Detail = b.Detail
	stepToUpdate.Status = b.Status
	stepUpdated := service.stepRepository.UpdateStep(stepToUpdate)

	updateStep := dto.StepUpdateDto{
		StepId: stepUpdated.Id,
		TodoId: uint64(stepUpdated.TodoId),
		Name:   stepUpdated.Name,
		Detail: stepUpdated.Detail,
		Status: stepUpdated.Status,
	}
	return updateStep
}

func (service *stepService) DeleteStepById(stepID uint64) error {
	err := service.stepRepository.DeleteStepById(stepID)
	if err != nil {
		return err
	}
	return nil
}

func (service *stepService) FindAllStep() []models.Step {
	return service.stepRepository.FindAllStep()
}

func (service *stepService) FindStepById(stepID uint64) models.Step {
	return service.stepRepository.FindStepById(stepID)
}
