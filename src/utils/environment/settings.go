package environment

import (
	"fmt"
	"microservice/src/configs"
	"os"

	"github.com/joho/godotenv"
)

func Get() {
	env := ".env"
	if err1 := godotenv.Load(env); err1 != nil {
		env = "../.env"
		if err2 := godotenv.Load(env); err2 != nil {
			fmt.Println(".env is not loaded properly")
			fmt.Println(err1)
			os.Exit(2)
		}
	}
	environment_select := configs.GetEnvironmentSelect()
	if err3 := godotenv.Load(fmt.Sprintf("%v.%v", env, environment_select)); err3 != nil {
		fmt.Printf(".env.%v is not loaded properly\n", environment_select)
		fmt.Println(err3)
		os.Exit(2)
	}
}
