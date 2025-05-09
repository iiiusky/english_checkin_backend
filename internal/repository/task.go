package repository

import (
	"context"
	"english_checkin_backend/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTask(ctx context.Context, id string) (*model.Task, error)
	ListTasks(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Task, int64, error)
	CreateTask(ctx context.Context, task *model.Task) error
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id string) error
}

func NewTaskRepository(
	repository *Repository,
	db *gorm.DB,
) TaskRepository {
	return &taskRepository{
		Repository: repository,
		db:         db,
	}
}

type taskRepository struct {
	*Repository
	db *gorm.DB
}

func (r *taskRepository) GetTask(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) ListTasks(ctx context.Context, page, pageSize int, query map[string]interface{}) ([]*model.Task, int64, error) {
	var tasks []*model.Task
	var total int64

	// 构建查询
	db := r.db.Model(&model.Task{})

	// 应用查询条件
	for key, value := range query {
		if key == "english_text" {
			db = db.Where("english_text LIKE ?", "%"+value.(string)+"%")
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
	err = db.Offset(offset).Limit(pageSize).Order("id DESC").Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

func (r *taskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(ctx context.Context, id string) error {
	return r.db.Delete(&model.Task{}, id).Error
}
