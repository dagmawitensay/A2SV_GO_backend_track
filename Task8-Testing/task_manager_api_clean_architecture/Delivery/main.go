package main

import (
	config "task_manager_api_clean_architecture/Config"
	"task_manager_api_clean_architecture/Delivery/routers"
	infrastructure "task_manager_api_clean_architecture/Infrastructure"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	configs := config.GetConfig()
	db := config.GetDB(configs)
	jwtService := infrastructure.NewJWTService([]byte(configs.SecretKey))

	r := gin.Default()
	
	routers.SetupRouter(10 * time.Second, db, r, jwtService)

	r.Run(":8080")
}