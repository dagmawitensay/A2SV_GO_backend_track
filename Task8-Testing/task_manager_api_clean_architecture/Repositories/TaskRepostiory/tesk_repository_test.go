package taskrepository

import (
	"context"
	config "task_manager_api_clean_architecture/Config"
	domain "task_manager_api_clean_architecture/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepositorySuite struct {
	suite.Suite
	repository domain.TaskRepository
	db         *mongo.Database
}

func (suite *taskRepositorySuite) SetupSuite() {
	configs, err := config.LoadConfig("../../.")
	suite.NoError(err)
	suite.db = config.GetDB(&configs)
	repository := NewTaskRepository(suite.db, "tasks")

	suite.repository = repository
}

func (suite *taskRepositorySuite) TearDownSuite() {
	err := suite.db.Drop(context.Background())
	suite.NoError(err)
}

func (suite *taskRepositorySuite) TearDownTest() {
	err := suite.db.Collection("tasks").Drop(context.Background())
	suite.NoError(err)

}

func (suite *taskRepositorySuite) TestCreateTask_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	err := suite.repository.Create(context.Background(), task)
	suite.NoError(err)

	var insertedTask domain.Task
	err = suite.db.Collection("tasks").FindOne(context.Background(), bson.M{"_id": task.ID}).Decode(&insertedTask)
	suite.NoError(err)
	suite.Equal(task.Title, insertedTask.Title)
	suite.Equal(task.Description, insertedTask.Description)
}

func (suite *taskRepositorySuite) TestGetAllTasks() {
	task1 := &domain.Task{
		Title:       "Test Task 1",
		Description: "Test Description 1",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	task2 := &domain.Task{
		Title:       "Test Task 2",
		Description: "Test Description 2",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      "completed",
	}

	_ = suite.repository.Create(context.Background(), task1)
	_ = suite.repository.Create(context.Background(), task2)

	tasks, err := suite.repository.GetAllTasks(context.Background())

	suite.NoError(err)
	suite.Len(tasks, 2)
}

func (suite *taskRepositorySuite) TestGetTaskById() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	_ = suite.repository.Create(context.Background(), task)

	retrievedTask, err := suite.repository.GetTaskById(context.Background(), task.ID.Hex())

	suite.NoError(err)
	suite.Equal(task.ID, retrievedTask.ID)
	suite.Equal(task.Title, retrievedTask.Title)
}

func (suite *taskRepositorySuite) TestTaskUpdate_Positive() {
	task := &domain.Task{
		Title:       "Original Task",
		Description: "Original Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	_ = suite.repository.Create(context.Background(), task)

	updateTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}

	err := suite.repository.UpdateTask(context.Background(), task.ID.Hex(), updateTask)

	suite.NoError(err)

	retrivedTask, err := suite.repository.GetTaskById(context.Background(), task.ID.Hex())
	suite.NoError(err)
	suite.Equal("Updated Task", retrivedTask.Title)
	suite.Equal("Updated Description", retrivedTask.Description)
	suite.Equal("completed", retrivedTask.Status)
}

func (suite *taskRepositorySuite) TestDeleteTask_Positive() {
	task := &domain.Task{
		Title:       "Task to be deleted",
		Description: "This task will be deleted",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	_ = suite.repository.Create(context.Background(), task)

	err := suite.repository.DeleteTask(context.Background(), task.ID.Hex())

	suite.NoError(err)

	retrivedTask, err := suite.repository.GetTaskById(context.Background(), task.ID.Hex())
	suite.Equal(mongo.ErrNoDocuments, err)
	suite.Nil(retrivedTask)
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(taskRepositorySuite))
}
