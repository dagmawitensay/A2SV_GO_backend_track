package taskusecases

import (
	"context"
	domain "task_manager_api_clean_architecture/Domain"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUseCase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUseCase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}


func (tu *taskUsecase) Create(c context.Context, task *domain.Task) error{
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) GetAllTasks(c context.Context)([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.GetAllTasks(ctx)
}

func (tu *taskUsecase) GetTaskById(c context.Context, id string)(*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.GetTaskById(ctx, id)
}

func (tu *taskUsecase) UpdateTask(c context.Context, id string, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.UpdateTask(ctx, id, task)
}

func (tu *taskUsecase) DeleteTask(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.DeleteTask(ctx, id)
}