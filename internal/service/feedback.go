package service

import (
    "context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type FeedbackService interface {
	GetFeedback(ctx context.Context, id int64) (*model.Feedback, error)
}
func NewFeedbackService(
    service *Service,
    feedbackRepository repository.FeedbackRepository,
) FeedbackService {
	return &feedbackService{
		Service:        service,
		feedbackRepository: feedbackRepository,
	}
}

type feedbackService struct {
	*Service
	feedbackRepository repository.FeedbackRepository
}

func (s *feedbackService) GetFeedback(ctx context.Context, id int64) (*model.Feedback, error) {
	return s.feedbackRepository.GetFeedback(ctx, id)
}
