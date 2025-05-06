package handler

import (
	"github.com/gin-gonic/gin"
	"english_checkin_backend/internal/service"
)

type ProfileHandler struct {
	*Handler
	profileService service.ProfileService
}

func NewProfileHandler(
    handler *Handler,
    profileService service.ProfileService,
) *ProfileHandler {
	return &ProfileHandler{
		Handler:      handler,
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(ctx *gin.Context) {

}
