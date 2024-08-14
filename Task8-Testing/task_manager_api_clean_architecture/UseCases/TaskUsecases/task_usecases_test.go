package taskusecases

import (
	"context"
	domain "task_manager_api_clean_architecture/Domain"
	"task_manager_api_clean_architecture/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type taskUsecaseSuite struct {
	suite.Suite
	repository *mocks.TaskRepository

	usecase domain.TaskUseCase
}

func (suite *taskUsecaseSuite) SetupTest() {
	repository := new(mocks.TaskRepository)
	usecase := NewTaskUseCase(repository, 10 * time.Second)

	suite.repository = repository
	suite.usecase = usecase
}

func (suite *taskUsecaseSuite) TestCreateTaks_Positive() {
	task := domain.Task{
		Title: "Test Task",
		Description: "Test Description",
		DueDate: time.Now().Add(24 * time.Hour),
		Status: "pending",
	}

	suite.repository.On("Create", mock.Anything, &task).Return(nil)

	err := suite.usecase.Create(context.Background(), &task)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetAllTasks() {
	tasks := []domain.Task{
		{
			Title:       "Test Task 1",
			Description: "This is test task 1",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Test Task 2",
			Description: "This is test task 2",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "completed",
		},
	}

	suite.repository.On("GetAllTasks", mock.Anything).Return(tasks, nil)

	result, err := suite.usecase.GetAllTasks(context.Background())

	suite.NoError(err)
	assert.Equal(suite.T(), tasks, result)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetTaskById() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	suite.repository.On("GetTaskById", mock.Anything, mock.AnythingOfType("string")).Return(task, nil)
	result, err := suite.usecase.GetTaskById(context.Background(), task.ID.Hex())
	suite.NoError(err)
	assert.Equal(suite.T(), task, result)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestUpdateTask() {
	task := &domain.Task{
		Title:       "Updated Task",
		Description: "This is an updated task",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      "completed",
	}

	suite.repository.On("UpdateTask", mock.Anything, mock.AnythingOfType("string"), task).Return(nil)
	err := suite.usecase.UpdateTask(context.Background(), task.ID.Hex(), task)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestDeleteTask() {
	taskID := "some-task-id"

	suite.repository.On("DeleteTask", mock.Anything, taskID).Return(nil)

	err := suite.usecase.DeleteTask(context.Background(), taskID)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(taskUsecaseSuite))
}