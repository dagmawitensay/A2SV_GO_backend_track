package taskrepository

import (
	"context"
	"errors"
	domain "task_manager_api_clean_architecture/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database *mongo.Database
	collection string
}


func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository{
	return &taskRepository{
		database: db,
		collection: collection,
	}
}


func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	if task.DueDate.Before(time.Now()) {
		return errors.New("due date can't be in the past")
	}

	collection := tr.database.Collection(tr.collection)

	insertOneResult, err := collection.InsertOne(c, task)

	if err != nil {
		return err
	}

	task.ID = insertOneResult.InsertedID.(primitive.ObjectID)

	return nil
}


func (tr *taskRepository) GetAllTasks(c context.Context)([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	
	var tasks  []domain.Task

	cursor, err := collection.Find(c, bson.D{{}})
	if err != nil {
		return nil, err
	} 

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}


func (tr *taskRepository) GetTaskById(c context.Context, id string)(*domain.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := tr.database.Collection(tr.collection)
	var task domain.Task

	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil

}


func (tr *taskRepository) UpdateTask(c context.Context, id string, updatedTask *domain.Task) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"_id": objectID}
	updatedFields := bson.M{}

	if updatedTask.Title != "" {
		updatedFields["title"] = updatedTask.Title
	}

	if updatedTask.Description != "" {
		updatedFields["description"] = updatedTask.Description
	}

	if !updatedTask.DueDate.IsZero() {
		updatedFields["due_date"] = updatedTask.DueDate
	}

	if updatedTask.Status != "" {
		updatedFields["status"] = updatedTask.Status
	}

	update := bson.M{
		"$set" : updatedFields,
	}

	_, err = collection.UpdateOne(context.Background(),filter, update)

	return err
}

func (tr *taskRepository) DeleteTask(c context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)

	return err
}