package repository

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"finpro/models"
)

type UserRepository interface {
	InsertUser(username string, password string, email string) models.User
	UpdateUser(userId uint64, username string, password string, email string) models.User
	VerifyCredential(username string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userId string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) InsertUser(username string, password string, email string) models.User {
	user := models.User{UserName: username, Password: password, UserGmail: email}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(userId uint64, username string, password string, email string) models.User {
	var user models.User
	db.connection.Find(&user, userId)
	user.UserName = username
	user.Password = password
	user.UserGmail = email
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user models.User
	res := db.connection.Where("user_name = ? AND password = ?", username, password).Find(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("user_gmail = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Find(&user, "user_gmail = ?", email)
	return user
}

func (db *userConnection) ProfileUser(userId string) models.User {
	var user models.User
	db.connection.Find(&user, "user_id = ?", userId)
	return user
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
