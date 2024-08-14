package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	domain "task_manager_api_clean_architecture/Domain"
	"task_manager_api_clean_architecture/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type controllerTestSuite struct {
	suite.Suite
	router *gin.Engine    
	taskUsecase *mocks.TaskUseCase
	userUsecase *mocks.UserUseCase
	controller *Controller
}

func (suite *controllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()

	suite.taskUsecase = new(mocks.TaskUseCase)
	suite.userUsecase = new(mocks.UserUseCase)
	suite.controller = &Controller{
		TaskUsecase: suite.taskUsecase,
		UserUsecase: suite.userUsecase,
	}

	suite.router.POST("/tasks", suite.controller.CreateTask)
	suite.router.GET("/tasks", suite.controller.GetAllTasks)
	suite.router.GET("/tasks/:id", suite.controller.GetTaskById)
	suite.router.PUT("/tasks/:id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.controller.DeleteTask)
	suite.router.POST("/register", suite.controller.Register)
	suite.router.POST("/login", suite.controller.Login)
	suite.router.PUT("/promote/:id", suite.controller.PromoteUser)
}

func (suite *controllerTestSuite) TestCreateTask_Positive() {

    suite.taskUsecase.On("Create", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil)

    body := `{"title":"Test Task","description":"Test Description","due_date":"2024-08-15T17:28:25Z","status":"pending"}`
    req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    suite.router.ServeHTTP(resp, req)


    assert.Equal(suite.T(), http.StatusCreated, resp.Code)
    assert.JSONEq(suite.T(), `{"message": "Task created successfully!"}`, resp.Body.String())
    suite.taskUsecase.AssertExpectations(suite.T())
}


func (suite *controllerTestSuite) TestGetAllTasks_Positive() {
	tasks := []domain.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	suite.taskUsecase.On("GetAllTasks", mock.Anything).Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	tasksJSON, err := json.Marshal(tasks)
    if err != nil {
        suite.T().Fatal(err)
    }


	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(),  string(tasksJSON), resp.Body.String())
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerTestSuite) TestGetTaskById_Positive() {
	task := domain.Task{
		Title: "Task",
	}

	suite.taskUsecase.On("GetTaskById", mock.Anything, "task-id").Return(&task, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/task-id", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	tasksJSON, err := json.Marshal(task)
    if err != nil {
        suite.T().Fatal(err)
    }

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), string(tasksJSON), resp.Body.String())
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerTestSuite) TestUpdateTask_Positive() {
	updatedTask := domain.Task{
		Title: "Updated Task",
	}

	suite.taskUsecase.On("UpdateTask", mock.Anything, "task-id", &updatedTask).Return(nil)

	body := `{"title":"Updated Task"}`
	req, _ := http.NewRequest(http.MethodPut, "/tasks/task-id", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), `{"message":"Task updated successfully!"}`, resp.Body.String())
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerTestSuite) TestDeleteTask_Positive() {
	suite.taskUsecase.On("DeleteTask", mock.Anything, mock.AnythingOfType("string")).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/task-id", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusNoContent, resp.Code)
	assert.Empty(suite.T(), resp.Body.String())
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerTestSuite) TestRegisterUser_Positive() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	suite.userUsecase.On("Create", mock.Anything, &user).Return(nil)

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusCreated, resp.Code)
	assert.JSONEq(suite.T(), `{"message": "User registerd successfully!"}`, resp.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *controllerTestSuite) TestLogin_Positive() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedToken := "mocked-jwt-token"

	suite.userUsecase.On("Login", mock.Anything, &user).Return(expectedToken, nil)

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), `{"message": "User logged in successfully!", "token": "mocked-jwt-token"}`, resp.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}


func TestController(t *testing.T) {
	suite.Run(t, new(controllerTestSuite))
}






