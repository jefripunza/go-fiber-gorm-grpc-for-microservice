package configs

import (
	"fmt"
	"os"
)

func DbType() string {
	return os.Getenv("DB_TYPE")
}
func DbHost() string {
	return os.Getenv("DB_HOST")
}
func DbPort() string {
	return os.Getenv("DB_PORT")
}
func DbUser() string {
	return os.Getenv("DB_USER")
}
func DbPass() string {
	return os.Getenv("DB_PASS")
}
func DbName() string {
	return os.Getenv("DB_NAME")
}

func DbLog() string {
	return os.Getenv("DB_LOG")
}
func DbSync() string {
	return os.Getenv("DB_SYNC")
}

func DatabaseConfig() string {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DbUser(), DbPass(), DbHost(), DbPort(), DbName())
}
