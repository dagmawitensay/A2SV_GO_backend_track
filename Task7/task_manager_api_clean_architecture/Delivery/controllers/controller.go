package controllers

import (
	"net/http"
	domain "task_manager_api_clean_architecture/task_manager_api_clean_architecture/Domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TaskUsecase domain.TaskUseCase
	UserUsecase domain.UserUseCase
}

func (ctr *Controller) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctr.TaskUsecase.Create(c, &task);
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task created successfully!"})
}

func (ctr *Controller) GetAllTasks(c *gin.Context) {
	task, err := ctr.TaskUsecase.GetAllTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (ctr *Controller) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := ctr.TaskUsecase.GetTaskById(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (ctr *Controller) UpdateTask(c *gin.Context) {
	var updatedTask domain.Task
	id := c.Param("id")

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	err := ctr.TaskUsecase.UpdateTask(c, id, &updatedTask);

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message" : "Task updated successfully!"})
}

func (ctr *Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := ctr.TaskUsecase.DeleteTask(c, id);
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Deleted Succesfully!"})
}

func (ctr *Controller) Register(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := ctr.UserUsecase.Create(c, &user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User registerd successfully!"})
}

func (ctr *Controller) Login(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	jwtToken, loginError := ctr.UserUsecase.Login(c, &user)
	if loginError != nil {
		c.IndentedJSON(http.StatusUnauthorized, loginError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "token": jwtToken})
}

func (ctr *Controller) PromoteUser(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.Promote(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully!"})
}