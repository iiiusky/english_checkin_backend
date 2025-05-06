package task

import (
	"english_checkin_backend/internal/repository"
	"english_checkin_backend/pkg/jwt"
	"english_checkin_backend/pkg/log"
	"english_checkin_backend/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
