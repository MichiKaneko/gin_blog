package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GinReleaseMode() bool {
	return os.Getenv("GIN_RELEASE_MODE") == "release"
}

func DatabaseConfig() string {
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	return db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
}
