package routers

import (
	"task_manager_api_clean_architecture/Delivery/controllers"
	domain "task_manager_api_clean_architecture/Domain"
	infrastructure "task_manager_api_clean_architecture/Infrastructure"
	repositories "task_manager_api_clean_architecture/Repositories"
	usecases "task_manager_api_clean_architecture/UseCases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	tr := repositories.NewTaskRepository(db, domain.CollectionTask)
	ur := repositories.NewUserRepository(db, domain.CollectionUser)
	ctr := &controllers.Controller{
		TaskUsecase: usecases.NewTaskUseCase(tr, timeout),
		UserUsecase: usecases.NewUserUseCase(ur, timeout),
	}
	
	publicRouter := gin.Group("")
	{
		publicRouter.POST("/register", ctr.Register)
		publicRouter.POST("/login", ctr.Login)
	}

	userRouter := gin.Group("")
	userRouter.Use(infrastructure.AuthMiddleWare())
	{
		userRouter.GET("/tasks", ctr.GetAllTasks)
		userRouter.GET("tasks/:id", ctr.GetTaskById)
	}

	adminRouter := gin.Group("")
	adminRouter.Use(infrastructure.AuthMiddleWare(), infrastructure.RoleMiddleware())
	{
		adminRouter.POST("/tasks", ctr.CreateTask)
		adminRouter.PUT("/tasks/:id", ctr.UpdateTask)
		adminRouter.DELETE("/tasks/:id", ctr.DeleteTask)
		adminRouter.PUT("/users/promote/:id", ctr.PromoteUser)
	}
}

