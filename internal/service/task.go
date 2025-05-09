package service

import (
	"context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type TaskService interface {
	GetTask(ctx context.Context, id string) (*model.Task, error)
	ListTasks(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Task, int64, error)
	CreateTask(ctx context.Context, task *model.Task) error
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id string) error
}

func NewTaskService(
	service *Service,
	taskRepository repository.TaskRepository,
) TaskService {
	return &taskService{
		Service:        service,
		taskRepository: taskRepository,
	}
}

type taskService struct {
	*Service
	taskRepository repository.TaskRepository
}

func (s *taskService) GetTask(ctx context.Context, id string) (*model.Task, error) {
	return s.taskRepository.GetTask(ctx, id)
}

func (s *taskService) ListTasks(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Task, int64, error) {
	return s.taskRepository.ListTasks(ctx, page, pageSize, query)
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) error {
	return s.taskRepository.CreateTask(ctx, task)
}

func (s *taskService) UpdateTask(ctx context.Context, task *model.Task) error {
	return s.taskRepository.UpdateTask(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepository.DeleteTask(ctx, id)
}
