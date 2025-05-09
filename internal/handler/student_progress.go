package handler

import (
	"english_checkin_backend/api/v1"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StudentProgressHandler struct {
	*Handler
	studentProgressService service.StudentProgressService
}

func NewStudentProgressHandler(
	handler *Handler,
	studentProgressService service.StudentProgressService,
) *StudentProgressHandler {
	return &StudentProgressHandler{
		Handler:                handler,
		studentProgressService: studentProgressService,
	}
}

// GetStudentProgress godoc
// @Summary 获取学生进度信息
// @Schemes
// @Description 根据学生ID和任务ID获取进度详细信息
// @Tags 学生进度模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param studentId query int true "学生ID"
// @Param taskId query int true "任务ID"
// @Success 200 {object} v1.GetStudentProgressResponse
// @Router /v1/english_check_in/student_progress [get]
func (h *StudentProgressHandler) GetStudentProgress(ctx *gin.Context) {
	var req v1.GetStudentProgressRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	progress, err := h.studentProgressService.GetStudentProgress(ctx, req.StudentId, req.TaskId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	data := v1.StudentProgressDataItem{
		ID:                  uint(progress.ID),
		StudentId:           progress.StudentId,
		TaskId:              progress.TaskId,
		CheckInAt:           progress.CheckInAt,
		AudioUrl:            progress.AudioUrl,
		AiScore:             progress.AiScore,
		PronAccuracyScore:   progress.PronAccuracyScore,
		PronFluencyScore:    progress.PronFluencyScore,
		PronCompletionScore: progress.PronCompletionScore,
		WordScores:          progress.WordScores,
		AiSuggestions:       progress.AiSuggestions,
		TeacherScore:        progress.TeacherScore,
		TeacherFeedback:     progress.TeacherFeedback,
		TeacherId:           progress.TeacherId,
		UpdatedAt:           progress.UpdatedAt.Format(time.RFC3339),
		CreatedAt:           progress.CreatedAt.Format(time.RFC3339),
	}

	v1.HandleSuccess(ctx, data)
}

// ListStudentProgress godoc
// @Summary 获取学生进度列表
// @Schemes
// @Description 分页获取学生进度列表信息
// @Tags 学生进度模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int true "页码"
// @Param pageSize query int true "每页数量"
// @Param studentId query int false "学生ID"
// @Param taskId query int false "任务ID"
// @Param startDate query int false "开始日期"
// @Param endDate query int false "结束日期"
// @Param taskType query string false "任务类型"
// @Param hasFeedback query bool false "是否已有教师评分"
// @Success 200 {object} v1.ListStudentProgressResponse
// @Router /v1/english_check_in/student_progress/list [get]
func (h *StudentProgressHandler) ListStudentProgress(ctx *gin.Context) {
	var req v1.ListStudentProgressRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}

	// 构建查询条件
	query := make(map[string]interface{})
	if req.StudentId != "" {
		query["student_id"] = req.StudentId
	}
	if req.TaskId != "" {
		query["task_id"] = req.TaskId
	}
	// 添加是否有教师评分的筛选条件
	if req.HasFeedback != nil {
		query["has_feedback"] = *req.HasFeedback
	}
	// 日期范围查询可以在repository层处理

	// 查询学生进度列表
	var progresses []*model.StudentProgress
	var total int64
	var err error
	
	// 查询学生进度列表
	progresses, total, err = h.studentProgressService.ListStudentProgress(ctx, req.Page, req.PageSize, query)
	
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	// 构建响应数据
	list := make([]v1.StudentProgressDataItem, 0, len(progresses))
	for _, progress := range progresses {
		list = append(list, v1.StudentProgressDataItem{
			ID:                  uint(progress.ID),
			StudentId:           progress.StudentId,
			TaskId:              progress.TaskId,
			CheckInAt:           progress.CheckInAt,
			AudioUrl:            progress.AudioUrl,
			AiScore:             progress.AiScore,
			PronAccuracyScore:   progress.PronAccuracyScore,
			PronFluencyScore:    progress.PronFluencyScore,
			PronCompletionScore: progress.PronCompletionScore,
			WordScores:          progress.WordScores,
			AiSuggestions:       progress.AiSuggestions,
			TeacherScore:        progress.TeacherScore,
			TeacherFeedback:     progress.TeacherFeedback,
			TeacherId:           progress.TeacherId,
			UpdatedAt:           progress.UpdatedAt.Format(time.RFC3339),
			CreatedAt:           progress.CreatedAt.Format(time.RFC3339),
			// 添加学生姓名
			StudentName:         progress.StudentName,
		})
	}

	data := v1.ListStudentProgressResponseData{
		List:  list,
		Total: total,
	}

	v1.HandleSuccess(ctx, data)
}

// GetStudentProgressById godoc
// @Summary 根据ID获取学生进度信息
// @Schemes
// @Description 根据进度ID获取详细信息，包含学生信息和任务信息
// @Tags 学生进度模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "进度ID"
// @Success 200 {object} v1.GetStudentProgressByIdResponse
// @Router /v1/english_check_in/student_progress/{id} [get]
func (h *StudentProgressHandler) GetStudentProgressById(ctx *gin.Context) {
	var req v1.GetStudentProgressByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	progress, err := h.studentProgressService.GetStudentProgressById(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	data := v1.StudentProgressDataItem{
		ID:                  uint(progress.ID),
		StudentId:           progress.StudentId,
		TaskId:              progress.TaskId,
		CheckInAt:           progress.CheckInAt,
		AudioUrl:            progress.AudioUrl,
		AiScore:             progress.AiScore,
		PronAccuracyScore:   progress.PronAccuracyScore,
		PronFluencyScore:    progress.PronFluencyScore,
		PronCompletionScore: progress.PronCompletionScore,
		WordScores:          progress.WordScores,
		AiSuggestions:       progress.AiSuggestions,
		TeacherScore:        progress.TeacherScore,
		TeacherFeedback:     progress.TeacherFeedback,
		TeacherId:           progress.TeacherId,
		UpdatedAt:           progress.UpdatedAt.Format(time.RFC3339),
		CreatedAt:           progress.CreatedAt.Format(time.RFC3339),
		// 添加关联数据
		StudentName:         progress.StudentName,
		TaskEnglishText:     progress.TaskEnglishText,
		TaskChineseText:     progress.TaskChineseText,
		TaskType:            progress.TaskType,
	}

	v1.HandleSuccess(ctx, data)
}
