package service

import (
	"context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type CheckInService interface {
	GetCheckIn(ctx context.Context, id int64) (*model.Task, error)
}

func NewCheckInService(
	service *Service,
	checkInRepository repository.DailyTaskRepository,
) CheckInService {
	return &checkInService{
		Service:           service,
		checkInRepository: checkInRepository,
	}
}

type checkInService struct {
	*Service
	checkInRepository repository.DailyTaskRepository
}

func (s *checkInService) GetCheckIn(ctx context.Context, id int64) (*model.Task, error) {
	return nil, nil
}
