package service

import (
	"context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type StudentProgressService interface {
	GetStudentProgress(ctx context.Context, studentId, taskId int64) (*model.StudentProgress, error)
	GetStudentProgressById(ctx context.Context, id int64) (*model.StudentProgress, error)
	ListStudentProgress(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.StudentProgress, int64, error)
}

func NewStudentProgressService(
	service *Service,
	studentProgressRepository repository.StudentProgressRepository,
) StudentProgressService {
	return &studentProgressService{
		Service:                   service,
		studentProgressRepository: studentProgressRepository,
	}
}

type studentProgressService struct {
	*Service
	studentProgressRepository repository.StudentProgressRepository
}

func (s *studentProgressService) GetStudentProgress(ctx context.Context, studentId, taskId int64) (*model.StudentProgress, error) {
	return s.studentProgressRepository.GetStudentProgress(ctx, studentId, taskId)
}

func (s *studentProgressService) GetStudentProgressById(ctx context.Context, id int64) (*model.StudentProgress, error) {
	return s.studentProgressRepository.GetStudentProgressById(ctx, id)
}

func (s *studentProgressService) ListStudentProgress(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.StudentProgress, int64, error) {
	return s.studentProgressRepository.ListStudentProgress(ctx, page, pageSize, query)
}
