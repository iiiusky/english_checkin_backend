package repository

import (
    "context"
	"english_checkin_backend/internal/model"
)

type FeedbackRepository interface {
	GetFeedback(ctx context.Context, id int64) (*model.Feedback, error)
}

func NewFeedbackRepository(
	repository *Repository,
) FeedbackRepository {
	return &feedbackRepository{
		Repository: repository,
	}
}

type feedbackRepository struct {
	*Repository
}

func (r *feedbackRepository) GetFeedback(ctx context.Context, id int64) (*model.Feedback, error) {
	var feedback model.Feedback

	return &feedback, nil
}
