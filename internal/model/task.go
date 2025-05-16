package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	Id              string `gorm:"column:id;type:varchar(100);primary_key;comment:Unique identifier for the task." json:"id"`
	EnglishText     string `gorm:"column:english_text;type:text;comment:The English text content of the task." json:"english_text"`
	ChineseText     string `gorm:"column:chinese_text;type:text;comment:The Chinese text content or translation for the task." json:"chinese_text"`
	StartTime       int64  `gorm:"column:start_time;type:bigint;comment:Timestamp indicating when the task becomes available or starts (milliseconds since epoch)." json:"start_time"`
	EndTime         int64  `gorm:"column:end_time;type:bigint;comment:Timestamp indicating when the task availability ends (milliseconds since epoch)." json:"end_time"`
	Students        string `gorm:"column:students;type:json;comment:Array of student IDs this task is assigned to." json:"students"`
	DifficultyLevel int    `gorm:"column:difficulty_level;type:int;comment:Difficulty level of the task (e.g., 1-5)." json:"difficulty_level"`
	TaskType        string `gorm:"column:task_type;type:varchar(100);comment:Type of task (e.g., 'reading', 'pronunciation')." json:"task_type"`
	CreatedAtMs     int64  `gorm:"column:createdAt;type:bigint;comment:Creation timestamp as milliseconds since epoch." json:"created_at_ms"`
	UpdatedAtMs     int64  `gorm:"column:updatedAt;type:bigint;comment:Last update timestamp as milliseconds since epoch." json:"updated_at_ms"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (m *Task) TableName() string {
	return "tasks"
}

func (m *Task) BeforeCreate(tx *gorm.DB) error {
	m.Id = uuid.New().String()
	m.CreatedAtMs = m.CreatedAt.UnixMilli()
	m.UpdatedAtMs = m.UpdatedAt.UnixMilli()
	return nil
}
