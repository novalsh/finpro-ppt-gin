package repository

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"finpro/models"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(UserGmail string, password string) interface{}
	IsDuplicateEmail(UserGmail string) (tx *gorm.DB)
	FindByEmail(UserGmail string) models.User
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

func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.UserId)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user models.User
	res := db.connection.Where("user_gmail = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(UserGmail string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("user_gmail = ?", UserGmail).Take(&user)
}

func (db *userConnection) FindByEmail(UserGmail string) models.User {
	var user models.User
	db.connection.Find(&user, "user_gmail = ?", UserGmail)
	return user
}

func (db *userConnection) ProfileUser(userId string) models.User {
	var user models.User
	db.connection.Find(&user, "user_id = ?", userId)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to has a password")
	}
	return string(hash)
}
