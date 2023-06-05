package service

import (
	"log"

	"github.com/mashingan/smapping"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type UserService interface {
	UpdateUser(b dto.UserUpdateDTO) models.User
	ProfileUser(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) UpdateUser(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) ProfileUser(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
