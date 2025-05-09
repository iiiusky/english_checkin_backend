package model

import "gorm.io/gorm"

type StudentProgress struct {
	gorm.Model
	StudentId           string  `gorm:"column:student_id"`
	TaskId              string  `gorm:"column:task_id"`
	CheckInAt           int     `gorm:"column:checkin_at"`
	AudioUrl            string  `gorm:"column:audio_url"`
	AiScore             float64 `gorm:"column:ai_score;default:0;comment:Overall score provided by the AI evaluation"`
	PronAccuracyScore   float64 `gorm:"column:pron_accuracy_score;default:0;comment:Pronunciation accuracy score from AI evaluation"`
	PronFluencyScore    float64 `gorm:"column:pron_fluency_score;default:0;comment:Pronunciation fluency score from AI evaluation"`
	PronCompletionScore float64 `gorm:"column:pron_completion_score;default:0;comment:Pronunciation completion score from AI evaluation"`
	WordScores          string  `gorm:"column:word_scores;type:json;comment:Detailed word-level scores and information from AI evaluation"`
	AiSuggestions       string  `gorm:"column:ai_suggestions;type:text;comment:Feedback and suggestions provided by the AI"`
	TeacherScore        float64 `gorm:"column:teacher_score;default:0;comment:Score provided by a human teacher, if applicable"`
	TeacherFeedback     string  `gorm:"column:teacher_feedback;type:text;comment:Feedback provided by a human teacher, if applicable"`
	TeacherId           string  `gorm:"column:teacher_id;type:varchar(255);comment:Foreign key referencing the teacher's id in an 'account' or 'teacher' table, if applicable"`
	
	// 关联字段，不存储在数据库中
	StudentName      string `gorm:"-"` // 学生姓名，从student表关联获取
	TaskEnglishText  string `gorm:"-"` // 任务英文内容，从task表关联获取
	TaskChineseText  string `gorm:"-"` // 任务中文内容，从task表关联获取
	TaskType         string `gorm:"-"` // 任务类型，从task表关联获取
}

func (m *StudentProgress) TableName() string {
	return "student_progress"
}
