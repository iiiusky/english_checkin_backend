package service

import (
	"context"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/repository"
)

type StudentService interface {
	GetStudent(ctx context.Context, id int64) (*model.Student, error)
	ListStudents(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Student, int64, error)
	UpdateStudent(ctx context.Context, student *model.Student) error
}

func NewStudentService(
	service *Service,
	studentRepository repository.StudentRepository,
) StudentService {
	return &studentService{
		Service:           service,
		studentRepository: studentRepository,
	}
}

type studentService struct {
	*Service
	studentRepository repository.StudentRepository
}

func (s *studentService) GetStudent(ctx context.Context, id int64) (*model.Student, error) {
	return s.studentRepository.GetStudent(ctx, id)
}

func (s *studentService) ListStudents(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Student, int64, error) {
	return s.studentRepository.ListStudents(ctx, page, pageSize, query)
}

func (s *studentService) UpdateStudent(ctx context.Context, student *model.Student) error {
	return s.studentRepository.UpdateStudent(ctx, student)
}
