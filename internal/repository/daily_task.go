package repository

import (
	"context"
	"english_checkin_backend/internal/model"
)

type DailyTaskRepository interface {
	GetDailyTask(ctx context.Context, id int64) (*model.Task, error)
}

func NewDailyTaskRepository(
	repository *Repository,
) DailyTaskRepository {
	return &dailyTaskRepository{
		Repository: repository,
	}
}

type dailyTaskRepository struct {
	*Repository
}

func (r *dailyTaskRepository) GetDailyTask(ctx context.Context, id int64) (*model.Task, error) {
	var dailyTask model.Task

	return &dailyTask, nil
}
