package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finpro/config"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the home page"})
	})

	r.Run()
}
