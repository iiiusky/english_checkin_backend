package repository

import (
	"context"
	"english_checkin_backend/internal/model"

	"gorm.io/gorm"
)

type StudentProgressRepository interface {
	GetStudentProgress(ctx context.Context, studentId, taskId int64) (*model.StudentProgress, error)
	GetStudentProgressById(ctx context.Context, id int64) (*model.StudentProgress, error)
	ListStudentProgress(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.StudentProgress, int64, error)
}

func NewStudentProgressRepository(
	repository *Repository,
	db *gorm.DB,
) StudentProgressRepository {
	return &studentProgressRepository{
		Repository: repository,
		db:         db,
	}
}

type studentProgressRepository struct {
	*Repository
	db *gorm.DB
}

func (r *studentProgressRepository) GetStudentProgress(ctx context.Context, studentId, taskId int64) (*model.StudentProgress, error) {
	var studentProgress model.StudentProgress

	err := r.db.Where("student_id = ? AND task_id = ?", studentId, taskId).First(&studentProgress).Error
	if err != nil {
		return nil, err
	}

	return &studentProgress, nil
}

func (r *studentProgressRepository) GetStudentProgressById(ctx context.Context, id int64) (*model.StudentProgress, error) {
	var studentProgress model.StudentProgress

	// 基本查询
	err := r.db.First(&studentProgress, id).Error
	if err != nil {
		return nil, err
	}

	// 查询学生信息
	var student model.Student
	err = r.db.Where("_openid = ?", studentProgress.StudentId).First(&student).Error
	if err == nil {
		// 如果找到学生信息，添加到进度记录中
		studentProgress.StudentName = student.Name
	}

	// 查询任务信息
	var task model.Task
	err = r.db.Where("id = ?", studentProgress.TaskId).First(&task).Error
	if err == nil {
		// 如果找到任务信息，添加到进度记录中
		studentProgress.TaskEnglishText = task.EnglishText
		studentProgress.TaskChineseText = task.ChineseText
		studentProgress.TaskType = task.TaskType
	}

	return &studentProgress, nil
}

func (r *studentProgressRepository) ListStudentProgress(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.StudentProgress, int64, error) {
	var studentProgresses []*model.StudentProgress
	var total int64
	
	// 构建查询
	db := r.db.Model(&model.StudentProgress{})
	
	// 应用查询条件
	for key, value := range query {
		// 处理是否有教师评分的特殊查询条件
		if key == "has_feedback" {
			hasFeedback, ok := value.(bool)
			if ok {
				if hasFeedback {
					// 已评分：teacher_score > 0 且 teacher_feedback 不为空
					db = db.Where("teacher_score > 0 AND teacher_feedback != ''")
				} else {
					// 未评分：teacher_score = 0 或 teacher_feedback 为空
					db = db.Where("teacher_score = 0 OR teacher_feedback = '' OR teacher_feedback IS NULL")
				}
			}
		} else {
			db = db.Where(key+" = ?", value)
		}
	}

	// 计算总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = db.Offset(offset).Limit(pageSize).Order("id DESC").Find(&studentProgresses).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询每个进度记录对应的学生信息
	for _, progress := range studentProgresses {
		// 查询学生信息
		var student model.Student
		err = r.db.Where("_openid = ?", progress.StudentId).First(&student).Error
		if err == nil {
			// 如果找到学生信息，添加到进度记录中
			progress.StudentName = student.Name
		}
	}

	return studentProgresses, total, nil
}
