package handler

import (
	"github.com/gin-gonic/gin"
	"english_checkin_backend/internal/service"
)

type CheckInHandler struct {
	*Handler
	checkInService service.CheckInService
}

func NewCheckInHandler(
    handler *Handler,
    checkInService service.CheckInService,
) *CheckInHandler {
	return &CheckInHandler{
		Handler:      handler,
		checkInService: checkInService,
	}
}

func (h *CheckInHandler) GetCheckIn(ctx *gin.Context) {

}
