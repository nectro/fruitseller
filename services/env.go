package services

import (
	"os"
)

func SetEnv() {
	os.Setenv("password","Samaresh@3396")
	os.Setenv("secretKey","TheSecretKey")
}
