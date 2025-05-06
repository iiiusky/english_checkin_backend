package service

import (
	"context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type DailyTaskService interface {
	GetDailyTask(ctx context.Context, id int64) (*model.Task, error)
}

func NewDailyTaskService(
	service *Service,
	dailyTaskRepository repository.DailyTaskRepository,
) DailyTaskService {
	return &dailyTaskService{
		Service:             service,
		dailyTaskRepository: dailyTaskRepository,
	}
}

type dailyTaskService struct {
	*Service
	dailyTaskRepository repository.DailyTaskRepository
}

func (s *dailyTaskService) GetDailyTask(ctx context.Context, id int64) (*model.Task, error) {
	return s.dailyTaskRepository.GetDailyTask(ctx, id)
}
