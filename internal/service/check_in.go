package service

import (
    "context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type CheckInService interface {
	GetCheckIn(ctx context.Context, id int64) (*model.CheckIn, error)
}
func NewCheckInService(
    service *Service,
    checkInRepository repository.CheckInRepository,
) CheckInService {
	return &checkInService{
		Service:        service,
		checkInRepository: checkInRepository,
	}
}

type checkInService struct {
	*Service
	checkInRepository repository.CheckInRepository
}

func (s *checkInService) GetCheckIn(ctx context.Context, id int64) (*model.CheckIn, error) {
	return s.checkInRepository.GetCheckIn(ctx, id)
}
