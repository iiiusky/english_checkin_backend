package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
}

func (m *Feedback) TableName() string {
    return "feedback"
}
