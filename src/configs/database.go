package configs

import (
	"fmt"
	"os"
)

var (
	DbType = os.Getenv("DB_TYPE")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")

	DbLog  = os.Getenv("DB_LOG")
	DbSync = os.Getenv("DB_SYNC")
)

func DatabaseConfig() string {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName)
}
