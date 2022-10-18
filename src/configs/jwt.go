package configs

import "os"

func GetJwtSecretKet() string {
	return os.Getenv("JWT_SECRET_KEY")
}
