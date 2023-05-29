package service

import (
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type AuthService interface {
	VerifyCredential(UserGmail string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(UserGmail string) models.User
}

type authService struct {
	userRepoistory repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepoistory: userRep,
	}
}

func (service *authService) VerifyCredential(UserGmail string, password string) interface{} {
	res := service.userRepoistory.VerifyCredential(UserGmail, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.UserGmail == UserGmail && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepoistory.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(UserGmail string) models.User {
	return service.userRepoistory.FindByEmail(UserGmail)
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
