package main

import (
	"gorm.io/gorm"

	"finpro/config"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

}
