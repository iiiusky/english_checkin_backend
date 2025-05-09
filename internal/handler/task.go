package handler

import (
	v1 "english_checkin_backend/api/v1"
	"english_checkin_backend/internal/model"
	"english_checkin_backend/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	*Handler
	taskService service.TaskService
}

func NewTaskHandler(
	handler *Handler,
	taskService service.TaskService,
) *TaskHandler {
	return &TaskHandler{
		Handler:     handler,
		taskService: taskService,
	}
}

// GetTask godoc
// @Summary 获取单个任务信息
// @Schemes
// @Description 根据ID获取任务详细信息
// @Tags 任务模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query int true "任务ID"
// @Success 200 {object} v1.GetTaskResponse
// @Router /v1/english_check_in/task [get]
func (h *TaskHandler) GetTask(ctx *gin.Context) {
	var req v1.GetTaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	task, err := h.taskService.GetTask(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	data := v1.TaskDataItem{
		ID:          task.Id,
		EnglishText: task.EnglishText,
		ChineseText: task.ChineseText,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,

		Students:        task.Students,
		DifficultyLevel: task.DifficultyLevel,
		TaskType:        task.TaskType,
		UpdatedAt:       task.UpdatedAt.Format(time.RFC3339),
		CreatedAt:       task.CreatedAt.Format(time.RFC3339),
	}

	v1.HandleSuccess(ctx, data)
}

// ListTasks godoc
// @Summary 获取任务列表
// @Schemes
// @Description 分页获取任务列表信息
// @Tags 任务模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int true "页码"
// @Param pageSize query int true "每页数量"
// @Param englishText query string false "英文内容"
// @Param taskType query string false "任务类型"
// @Param difficultyLevel query int false "难度等级"
// @Success 200 {object} v1.GetTasksResponse
// @Router /v1/english_check_in/tasks [get]
func (h *TaskHandler) ListTasks(ctx *gin.Context) {
	var req v1.GetTasksRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}

	// 构建查询条件
	query := make(map[string]interface{})
	if req.EnglishText != "" {
		query["english_text"] = req.EnglishText
	}
	if req.TaskType != "" {
		query["task_type"] = req.TaskType
	}
	if req.DifficultyLevel > 0 {
		query["difficulty_level"] = req.DifficultyLevel
	}

	// 查询任务列表
	tasks, total, err := h.taskService.ListTasks(ctx, req.Page, req.PageSize, query)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	// 构建响应数据
	list := make([]v1.TaskDataItem, 0, len(tasks))
	for _, task := range tasks {
		list = append(list, v1.TaskDataItem{
			ID:          task.Id,
			EnglishText: task.EnglishText,
			ChineseText: task.ChineseText,
			StartTime:   task.StartTime,
			EndTime:     task.EndTime,

			Students:        task.Students,
			DifficultyLevel: task.DifficultyLevel,
			TaskType:        task.TaskType,
			UpdatedAt:       task.UpdatedAt.Format(time.RFC3339),
			CreatedAt:       task.CreatedAt.Format(time.RFC3339),
		})
	}

	data := v1.GetTasksResponseData{
		List:  list,
		Total: total,
	}

	v1.HandleSuccess(ctx, data)
}

// UpdateTask godoc
// @Summary 更新任务信息
// @Schemes
// @Description 更新指定任务的信息
// @Tags 任务模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param task body v1.TaskUpdateRequest true "任务信息"
// @Success 200 {object} v1.Response
// @Router /v1/english_check_in/task [put]
func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	var req v1.TaskUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 先获取任务信息
	task, err := h.taskService.GetTask(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	// 更新任务信息
	task.EnglishText = req.EnglishText
	task.ChineseText = req.ChineseText
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime

	task.Students = req.Students
	task.DifficultyLevel = req.DifficultyLevel
	task.TaskType = req.TaskType

	// 保存更新
	err = h.taskService.UpdateTask(ctx, task)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// CreateTask godoc
// @Summary 创建新任务
// @Schemes
// @Description 创建一个新的任务
// @Tags 任务模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param task body v1.TaskCreateRequest true "任务信息"
// @Success 200 {object} v1.Response
// @Router /v1/english_check_in/task [post]
func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var req v1.TaskCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 创建新任务
	task := &model.Task{
		EnglishText: req.EnglishText,
		ChineseText: req.ChineseText,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,

		Students:        req.Students,
		DifficultyLevel: req.DifficultyLevel,
		TaskType:        req.TaskType,
	}

	// 保存任务
	err := h.taskService.CreateTask(ctx, task)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteTask godoc
// @Summary 删除任务
// @Schemes
// @Description 删除指定的任务
// @Tags 任务模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query int true "任务ID"
// @Success 200 {object} v1.Response
// @Router /v1/english_check_in/task [delete]
func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	var req v1.TaskDeleteRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 删除任务
	err := h.taskService.DeleteTask(ctx, req.ID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
