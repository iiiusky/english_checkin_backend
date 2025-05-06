package job

import (
	"english_checkin_backend/internal/repository"
	"english_checkin_backend/pkg/jwt"
	"english_checkin_backend/pkg/log"
	"english_checkin_backend/pkg/sid"
)

type Job struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
