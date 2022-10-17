package apps

import (
	"fmt"
	"log"
	"os"
	"time"

	"main-service/src/configs"
	"main-service/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	db_type := os.Getenv("DB_TYPE")
	db_log := os.Getenv("DB_LOG")
	var connection *gorm.DB
	var err_msg error
	dsn := configs.DatabaseConfig()
	var db_config *gorm.Config
	if db_log == "true" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
		db_config = &gorm.Config{
			Logger: newLogger,
		}
	} else {
		db_config = &gorm.Config{}
	}
	if db_type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%v.db", os.Getenv("DB_NAME"))), db_config)
		connection = db
		err_msg = err
	} else if db_type == "mysql" {
		db, err := gorm.Open(mysql.Open(dsn), db_config)
		connection = db
		err_msg = err
	} else if db_type == "postgres" {
		db, err := gorm.Open(postgres.Open(dsn), db_config)
		connection = db
		err_msg = err
	} else {
		panic("hey dude, database type is not found!")
	}
	if err_msg != nil {
		log.Fatalln(err_msg)
		panic("failed to Database database")
	}
	return connection
}

func DatabaseMigration() {
	db_sync := os.Getenv("DB_SYNC")
	if db_sync == "true" {
		fmt.Println("Database Migrating...")
		db := DatabaseConnect()
		models.MigrationTables(db)
		fmt.Println("Migrate Success!")
	} else {
		fmt.Println("Database Off Migrating!")
	}
}
