package apps

import (
	"fmt"
	"log"
	"os"
	"time"

	"microservice/src/configs"
	"microservice/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {

	db_log := configs.DbLog()
	db_type := configs.DbType()
	dsn := configs.DatabaseConfig()

	var db_config *gorm.Config
	if db_log == "true" {
		db_config = &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second,   // Slow SQL threshold
					LogLevel:                  logger.Silent, // Log level
					IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
					Colorful:                  false,         // Disable color
				},
			),
		}
	} else {
		db_config = &gorm.Config{}
	}

	var connection *gorm.DB
	var err_msg error
	if db_type == "sqlite" {
		connection, err_msg = gorm.Open(sqlite.Open(fmt.Sprintf("%v.db", configs.DbName())), db_config)
	} else if db_type == "mysql" {
		connection, err_msg = gorm.Open(mysql.Open(dsn), db_config)
	} else if db_type == "postgres" {
		connection, err_msg = gorm.Open(postgres.Open(dsn), db_config)
	} else {
		panic("hey dude, database type is not found!")
	}

	if err_msg != nil {
		log.Fatalln(err_msg)
		panic("failed to Database database")
	}

	if db_log == "true" {
		return connection.Debug()
	} else {
		return connection
	}

}

func DatabaseMigration() {

	go func() {
		time.Sleep(22)
		if configs.DbSync() == "true" {
			fmt.Println("Database Migrating...")
			models.MigrationTables(DatabaseConnect())
			fmt.Println("Database Migrate Success!")
		} else {
			fmt.Println("Database Off Migrating!")
		}
	}()

}
