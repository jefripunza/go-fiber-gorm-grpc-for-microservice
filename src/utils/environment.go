package utils

import (
	"fmt"
	"main-service/src/configs"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironment() {
	if err1 := godotenv.Load(".env"); err1 != nil {
		fmt.Println(".env is not loaded properly")
		fmt.Println(err1)
		os.Exit(2)
	} else {
		environment_select := configs.GetEnvironmentSelect()
		if err2 := godotenv.Load(fmt.Sprintf(".env.%v", environment_select)); err2 != nil {
			fmt.Printf(".env.%v is not loaded properly\n", environment_select)
			fmt.Println(err2)
			os.Exit(2)
		}
	}

}
