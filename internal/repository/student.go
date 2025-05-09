package repository

import (
    "context"
	"english_checkin_backend/internal/model"
)

type StudentRepository interface {
	GetStudent(ctx context.Context, id int64) (*model.Student, error)
	ListStudents(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Student, int64, error)
	UpdateStudent(ctx context.Context, student *model.Student) error
	CreateStudent(ctx context.Context, student *model.Student) error
	DeleteStudent(ctx context.Context, id int64) error
}

func NewStudentRepository(
	repository *Repository,
) StudentRepository {
	return &studentRepository{
		Repository: repository,
	}
}

type studentRepository struct {
	*Repository
}

func (r *studentRepository) GetStudent(ctx context.Context, id int64) (*model.Student, error) {
	var student model.Student
	result := r.DB(ctx).Where("id = ?", id).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (r *studentRepository) ListStudents(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Student, int64, error) {
	var students []*model.Student
	var total int64

	db := r.DB(ctx).Model(&model.Student{})
	
	// Apply filters
	if name, ok := query["name"].(string); ok && name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	
	if phoneNumber, ok := query["phone_number"].(string); ok && phoneNumber != "" {
		db = db.Where("phone_number LIKE ?", "%"+phoneNumber+"%")
	}
	
	if status, ok := query["status"].(string); ok && status != "" {
		db = db.Where("status = ?", status)
	}
	
	// Count total records
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Pagination
	offset := (page - 1) * pageSize
	result := db.Offset(offset).Limit(pageSize).Find(&students)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	
	return students, total, nil
}

func (r *studentRepository) UpdateStudent(ctx context.Context, student *model.Student) error {
	// Use Updates instead of Save to only update non-zero fields
	// This prevents issues with zero datetime values
	result := r.DB(ctx).Model(&model.Student{}).Where("id = ?", student.ID).Updates(map[string]interface{}{
		"name":         student.Name,
		"avatar_url":   student.AvatarURL,
		"phone_number": student.PhoneNumber,
		"expiry_date":  student.ExpiryDate,
		"status":       student.Status,
		"_openid":      student.OpenID,
	})
	return result.Error
}

func (r *studentRepository) CreateStudent(ctx context.Context, student *model.Student) error {
	result := r.DB(ctx).Create(student)
	return result.Error
}

func (r *studentRepository) DeleteStudent(ctx context.Context, id int64) error {
	result := r.DB(ctx).Delete(&model.Student{}, id)
	return result.Error
}
