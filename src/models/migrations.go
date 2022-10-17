package models

import (
	"main-service/src/models/entities"

	"gorm.io/gorm"
)

func MigrationTables(db *gorm.DB) {

	// Register Your Table Here...
	db.AutoMigrate(&entities.Products{})

}
