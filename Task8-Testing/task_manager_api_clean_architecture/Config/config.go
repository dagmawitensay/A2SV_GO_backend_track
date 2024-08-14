package config

import (
	"fmt"
	"log"
	"os"
	domain "task_manager_api_clean_architecture/Domain"

	"github.com/joho/godotenv"
)

func GetConfig() *domain.Config {
	wd, err := os.Getwd()
	fmt.Println(wd)

	err =  godotenv.Load()

	if err != nil {
		log.Fatal("Error while reading env file: ", err)
	}
	
	config := &domain.Config{
		Database: domain.DatabaseConfig{
			DBURI:    os.Getenv("DB_URI"),
			DbName:   os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		TimeZone:  "UTC/GMT +3",
		SecretKey: os.Getenv("JWT_SECRET"),
	}
	
	return config
	
}

	 