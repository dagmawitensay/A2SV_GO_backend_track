package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)


type Task struct {
	ID  			primitive.ObjectID 		`bson:"_id,omitempty" json:"id,omitempty"`
	Title  			string 					`bson:"title" json:"title"`
	Description 	string 					`bson:"description" json:"description"`
	DueDate 		time.Time 				`bson:"due_date" json:"due_date"`
	Status 			string 					`bson:"status" json:"status"`
}


type TaskRepository interface {
	Create(c context.Context, task *Task) error
	GetAllTasks(c context.Context)([]Task, error)
	GetTaskById(c context.Context, id string)(*Task, error)
	UpdateTask(c context.Context, id string, task *Task) error
	DeleteTask(c context.Context, id string) error
}

type TaskUseCase interface {
	Create(c context.Context, task *Task) error
	GetAllTasks(c context.Context)([]Task, error)
	GetTaskById(c context.Context, id string)(*Task, error)
	UpdateTask(c context.Context, id string, task *Task) error
	DeleteTask(c context.Context, id string) error
}