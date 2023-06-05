package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type AuthService interface {
	VerifyCredential(Email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(Email string) models.User
	IsDuplicateEmail(Email string) bool
}

type authService struct {
	userRepoistory repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepoistory: userRep,
	}
}

func (service *authService) VerifyCredential(Email string, password string) interface{} {
	res := service.userRepoistory.VerifyCredential(Email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == Email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	userToCreate.Name = user.Name
	userToCreate.Email = user.Email
	userToCreate.Phone = user.Phone
	userToCreate.Password = user.Password

	InsertToDB := service.userRepoistory.InsertUser(userToCreate)

	return InsertToDB
}

func (service *authService) FindByEmail(Email string) models.User {
	return service.userRepoistory.FindByEmail(Email)
}

func (service *authService) IsDuplicateEmail(Email string) bool {
	res := service.userRepoistory.IsDuplicateEmail(Email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
