package handler

import (
	"english_checkin_backend/api/v1"
	"english_checkin_backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type StudentHandler struct {
	*Handler
	studentService service.StudentService
}

func NewStudentHandler(
	handler *Handler,
	studentService service.StudentService,
) *StudentHandler {
	return &StudentHandler{
		Handler:        handler,
		studentService: studentService,
	}
}

// GetStudent godoc
// @Summary 获取单个学生信息
// @Schemes
// @Description 根据ID获取学生详细信息
// @Tags 学生模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query int true "学生ID"
// @Success 200 {object} v1.GetStudentResponse
// @Router /v1/english_check_in/student [get]
func (h *StudentHandler) GetStudent(ctx *gin.Context) {
	var req v1.GetStudentRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	student, err := h.studentService.GetStudent(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	data := v1.StudentDataItem{
		ID:          uint(student.ID),
		OpenID:      student.OpenID,
		Name:        student.Name,
		AvatarURL:   student.AvatarURL,
		PhoneNumber: student.PhoneNumber,
		ExpiryDate:  student.ExpiryDate,
		Status:      student.Status,
		UpdatedAt:   student.UpdatedAt.Format(time.RFC3339),
		CreatedAt:   student.CreatedAt.Format(time.RFC3339),
	}

	v1.HandleSuccess(ctx, data)
}

// ListStudents godoc
// @Summary 获取学生列表
// @Schemes
// @Description 分页获取学生列表信息
// @Tags 学生模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int true "页码"
// @Param pageSize query int true "每页数量"
// @Param name query string false "学生姓名"
// @Param phoneNumber query string false "手机号"
// @Param status query string false "状态"
// @Success 200 {object} v1.GetStudentsResponse
// @Router /v1/english_check_in/students [get]
func (h *StudentHandler) ListStudents(ctx *gin.Context) {
	var req v1.GetStudentsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}

	// 构建查询条件
	query := make(map[string]interface{})
	if req.Name != "" {
		query["name"] = req.Name
	}
	if req.PhoneNumber != "" {
		query["phone_number"] = req.PhoneNumber
	}
	if req.Status != "" {
		query["status"] = req.Status
	}

	// 查询学生列表
	students, total, err := h.studentService.ListStudents(ctx, req.Page, req.PageSize, query)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	// 构建响应数据
	list := make([]v1.StudentDataItem, 0, len(students))
	for _, student := range students {
		list = append(list, v1.StudentDataItem{
			ID:          uint(student.ID),
			OpenID:      student.OpenID,
			Name:        student.Name,
			AvatarURL:   student.AvatarURL,
			PhoneNumber: student.PhoneNumber,
			ExpiryDate:  student.ExpiryDate,
			Status:      student.Status,
			UpdatedAt:   student.UpdatedAt.Format(time.RFC3339),
			CreatedAt:   student.CreatedAt.Format(time.RFC3339),
		})
	}

	data := v1.GetStudentsResponseData{
		List:  list,
		Total: total,
	}

	v1.HandleSuccess(ctx, data)
}

// UpdateStudent godoc
// @Summary 更新学生信息
// @Schemes
// @Description 更新指定学生的信息
// @Tags 学生模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param student body v1.StudentUpdateRequest true "学生信息"
// @Success 200 {object} v1.Response
// @Router /v1/english_check_in/student [put]
func (h *StudentHandler) UpdateStudent(ctx *gin.Context) {
	var req v1.StudentUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 先获取学生信息
	student, err := h.studentService.GetStudent(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	// 更新学生信息
	student.Name = req.Name
	student.AvatarURL = req.AvatarURL
	student.PhoneNumber = req.PhoneNumber
	student.ExpiryDate = req.ExpiryDate
	student.Status = req.Status

	// 保存更新
	err = h.studentService.UpdateStudent(ctx, student)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
