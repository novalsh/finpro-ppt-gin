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
	VerifyCredential(Email string, password string) interface{}
	IsDuplicateEmail(Email string) (tx *gorm.DB)
	FindByEmail(Email string) models.User
	ProfileUser(ID string) models.User
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
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(Email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", Email).Take(&user)
}

func (db *userConnection) FindByEmail(Email string) models.User {
	var user models.User
	db.connection.Find(&user, "email = ?", Email)
	return user
}

func (db *userConnection) ProfileUser(ID string) models.User {
	var user models.User
	db.connection.Find(&user, "id = ?", ID)
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
