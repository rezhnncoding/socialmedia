package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionString = ""
	Port             = 0
	SecretKey        []byte
)

func Load() {
	fmt.Println("Environment variables load started.")

	error := godotenv.Load()
	if error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if error != nil {
		log.Fatal(error)
	}

	ConnectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	fmt.Println("Environment variables load completed.")
}
