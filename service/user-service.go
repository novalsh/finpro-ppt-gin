package service

import (
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

func (service *userService) UpdateUser(b dto.UserUpdateDTO) models.User {
	user := models.User{}
	user.ID = b.Id
	user.Name = b.Name
	user.Email = b.Email
	user.Password = b.Password
	res := service.userRepository.UpdateUser(user)
	return res
}

func (service *userService) ProfileUser(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
