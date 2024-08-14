package main

import (
	config "task_manager_api_clean_architecture/Config"
	"task_manager_api_clean_architecture/Delivery/routers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	configs := config.GetConfig()
	db := config.GetDB(configs)

	r := gin.Default()
	
	routers.SetupRouter(10 * time.Second, db, r)

	r.Run(":8080")
}