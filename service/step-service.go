package service

import (
	"log"

	"github.com/mashingan/smapping"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type StepService interface {
	InsertStep(b dto.StepCreateDto) models.Step
	UpdateStep(b dto.StepUpdateDto) models.Step
	DeleteStep(b models.Step)
	FindAllStepAll() []models.Step
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

func (service *stepService) InsertStep(b dto.StepCreateDto) models.Step {
	step := models.Step{}
	err := smapping.FillStruct(&step, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.stepRepository.InsertStep(step)
	return res
}

func (service *stepService) UpdateStep(b dto.StepUpdateDto) models.Step {
	step := models.Step{}
	err := smapping.FillStruct(&step, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.stepRepository.UpdateStep(step)
	return res
}

func (service *stepService) DeleteStep(b models.Step) {
	service.stepRepository.DeleteStep(b)
}

func (service *stepService) FindAllStepAll() []models.Step {
	return service.stepRepository.FindAllStep()
}

func (service *stepService) FindStepById(stepID uint64) models.Step {
	return service.stepRepository.FindStepById(stepID)
}
