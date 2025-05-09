package v1

// Task data structure for API responses
type TaskDataItem struct {
	ID              string `json:"id"`
	EnglishText     string `json:"english_text"`
	ChineseText     string `json:"chinese_text"`
	StartTime       int64  `json:"start_time"`
	EndTime         int64  `json:"end_time"`
	Students        string `json:"students"`
	DifficultyLevel int    `json:"difficulty_level"`
	TaskType        string `json:"task_type"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}

// GetTaskRequest - Request to get a single task by ID
type GetTaskRequest struct {
	ID string `form:"id" binding:"required" example:"1"`
}

// GetTaskResponse - Response for a single task
type GetTaskResponse struct {
	Response
	Data TaskDataItem `json:"data"`
}

// GetTasksRequest - Request for paginated task list
type GetTasksRequest struct {
	Page            int    `form:"page" binding:"" example:"1"`
	PageSize        int    `form:"pageSize" binding:"required" example:"10"`
	EnglishText     string `form:"englishText" binding:"" example:"Hello world"`
	TaskType        string `form:"taskType" binding:"" example:"reading"`
	DifficultyLevel int    `form:"difficultyLevel" binding:"" example:"3"`
}

// GetTasksResponseData - Data for paginated task list
type GetTasksResponseData struct {
	List  []TaskDataItem `json:"list"`
	Total int64          `json:"total"`
}

// GetTasksResponse - Response for paginated task list
type GetTasksResponse struct {
	Response
	Data GetTasksResponseData `json:"data"`
}

// TaskCreateRequest - Request to create a new task
type TaskCreateRequest struct {
	EnglishText     string `json:"english_text" binding:"required" example:"Hello world"`
	ChineseText     string `json:"chinese_text" binding:"" example:"你好，世界"`
	StartTime       int64  `json:"start_time" binding:"required" example:"1714608000000"`
	EndTime         int64  `json:"end_time" binding:"required" example:"1714694400000"`
	Students        string `json:"students" binding:"" example:"[1,2,3]"`
	DifficultyLevel int    `json:"difficulty_level" binding:"required" example:"3"`
	TaskType        string `json:"task_type" binding:"required" example:"reading"`
}

// TaskUpdateRequest - Request to update an existing task
type TaskUpdateRequest struct {
	ID              string `json:"id" binding:"required" example:"1"`
	EnglishText     string `json:"english_text" binding:"required" example:"Hello world"`
	ChineseText     string `json:"chinese_text" binding:"" example:"你好，世界"`
	StartTime       int64  `json:"start_time" binding:"required" example:"1714608000000"`
	EndTime         int64  `json:"end_time" binding:"required" example:"1714694400000"`
	Students        string `json:"students" binding:"" example:"[1,2,3]"`
	DifficultyLevel int    `json:"difficulty_level" binding:"required" example:"3"`
	TaskType        string `json:"task_type" binding:"required" example:"reading"`
}

// TaskDeleteRequest - Request to delete a task
type TaskDeleteRequest struct {
	ID string `form:"id" binding:"required" example:"1"`
}
