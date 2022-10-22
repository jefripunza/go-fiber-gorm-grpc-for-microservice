package models

import (
	"microservice/src/models/entities"

	"gorm.io/gorm"
)

func MigrationTables(db *gorm.DB) {

	// Register Your Table Here...
	db.AutoMigrate(&entities.Products{})

}
