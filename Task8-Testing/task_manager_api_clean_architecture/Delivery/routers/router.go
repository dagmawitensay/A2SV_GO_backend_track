package routers

import (
	"task_manager_api_clean_architecture/Delivery/controllers"
	domain "task_manager_api_clean_architecture/Domain"
	infrastructure "task_manager_api_clean_architecture/Infrastructure"
	taskrepository "task_manager_api_clean_architecture/Repositories/TaskRepostiory"
	userrepository "task_manager_api_clean_architecture/Repositories/UserRepository"
	taskusecases "task_manager_api_clean_architecture/UseCases/TaskUsecases"
	userusecases "task_manager_api_clean_architecture/UseCases/UserUsecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	tr := taskrepository.NewTaskRepository(db, domain.CollectionTask)
	ur := userrepository.NewUserRepository(db, domain.CollectionUser)
	
	ctr := &controllers.Controller{
		TaskUsecase: taskusecases.NewTaskUseCase(tr, timeout),
		UserUsecase: userusecases.NewUserUseCase(ur, timeout),
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

