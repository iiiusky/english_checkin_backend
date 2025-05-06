package handler

import (
	"github.com/gin-gonic/gin"
	"english_checkin_backend/internal/service"
)

type FeedbackHandler struct {
	*Handler
	feedbackService service.FeedbackService
}

func NewFeedbackHandler(
    handler *Handler,
    feedbackService service.FeedbackService,
) *FeedbackHandler {
	return &FeedbackHandler{
		Handler:      handler,
		feedbackService: feedbackService,
	}
}

func (h *FeedbackHandler) GetFeedback(ctx *gin.Context) {

}
