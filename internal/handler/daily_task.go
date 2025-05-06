package handler

import (
	"github.com/gin-gonic/gin"
	"english_checkin_backend/internal/service"
)

type DailyTaskHandler struct {
	*Handler
	dailyTaskService service.DailyTaskService
}

func NewDailyTaskHandler(
    handler *Handler,
    dailyTaskService service.DailyTaskService,
) *DailyTaskHandler {
	return &DailyTaskHandler{
		Handler:      handler,
		dailyTaskService: dailyTaskService,
	}
}

func (h *DailyTaskHandler) GetDailyTask(ctx *gin.Context) {

}
