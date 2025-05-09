package v1

// StudentProgress data structure for API responses
type StudentProgressDataItem struct {
	ID                  uint    `json:"id"`
	StudentId           string  `json:"student_id"`
	TaskId              string  `json:"task_id"`
	CheckInAt           int     `json:"checkin_at"`
	AudioUrl            string  `json:"audio_url"`
	AiScore             float64 `json:"ai_score"`
	PronAccuracyScore   float64 `json:"pron_accuracy_score"`
	PronFluencyScore    float64 `json:"pron_fluency_score"`
	PronCompletionScore float64 `json:"pron_completion_score"`
	WordScores          string  `json:"word_scores"`
	AiSuggestions       string  `json:"ai_suggestions"`
	TeacherScore        float64 `json:"teacher_score"`
	TeacherFeedback     string  `json:"teacher_feedback"`
	TeacherId           string  `json:"teacher_id"`
	UpdatedAt           string  `json:"updatedAt"`
	CreatedAt           string  `json:"createdAt"`
	// Additional fields for joined data
	StudentName     string `json:"student_name,omitempty"`
	TaskEnglishText string `json:"task_english_text,omitempty"`
	TaskChineseText string `json:"task_chinese_text,omitempty"`
	TaskType        string `json:"task_type,omitempty"`
}

// GetStudentProgressRequest - Request to get student progress by student ID and task ID
type GetStudentProgressRequest struct {
	StudentId int64 `form:"studentId" binding:"required" example:"1"`
	TaskId    int64 `form:"taskId" binding:"required" example:"1"`
}

// GetStudentProgressResponse - Response for student progress
type GetStudentProgressResponse struct {
	Response
	Data StudentProgressDataItem `json:"data"`
}

// GetStudentProgressByIdRequest - Request to get student progress by ID
type GetStudentProgressByIdRequest struct {
	ID int64 `uri:"id" binding:"required" example:"1"`
}

// GetStudentProgressByIdResponse - Response for student progress by ID
type GetStudentProgressByIdResponse struct {
	Response
	Data StudentProgressDataItem `json:"data"`
}

// ListStudentProgressRequest - Request for paginated student progress list
type ListStudentProgressRequest struct {
	Page        int    `form:"page" binding:"" example:"1"`
	PageSize    int    `form:"pageSize" binding:"required" example:"10"`
	StudentId   string `form:"studentId" binding:"" example:"oU3iR6-qEHGdNcc8RMcOSBHoJFZg"`
	TaskId      string `form:"taskId" binding:"" example:"task-123"`
	StartDate   int64  `form:"startDate" binding:"" example:"1714608000000"`
	EndDate     int64  `form:"endDate" binding:"" example:"1714694400000"`
	TaskType    string `form:"taskType" binding:"" example:"reading"`
	HasFeedback *bool  `form:"hasFeedback" binding:"" example:"false"`
}

// ListStudentProgressResponseData - Data for paginated student progress list
type ListStudentProgressResponseData struct {
	List  []StudentProgressDataItem `json:"list"`
	Total int64                     `json:"total"`
}

// ListStudentProgressResponse - Response for paginated student progress list
type ListStudentProgressResponse struct {
	Response
	Data ListStudentProgressResponseData `json:"data"`
}

// StudentProgressCreateRequest - Request to create a new student progress
type StudentProgressCreateRequest struct {
	StudentId           int     `json:"student_id" binding:"required" example:"1"`
	TaskId              int     `json:"task_id" binding:"required" example:"1"`
	CheckInAt           int     `json:"checkin_at" binding:"required" example:"1714608000"`
	AudioUrl            string  `json:"audio_url" binding:"" example:"https://example.com/audio.mp3"`
	AiScore             float64 `json:"ai_score" binding:"" example:"85.5"`
	PronAccuracyScore   float64 `json:"pron_accuracy_score" binding:"" example:"80.0"`
	PronFluencyScore    float64 `json:"pron_fluency_score" binding:"" example:"90.0"`
	PronCompletionScore float64 `json:"pron_completion_score" binding:"" example:"85.0"`
	WordScores          string  `json:"word_scores" binding:"" example:"[{\"word\":\"hello\",\"score\":90}]"`
	AiSuggestions       string  `json:"ai_suggestions" binding:"" example:"Work on the pronunciation of 'th' sounds."`
}

// StudentProgressUpdateRequest - Request to update an existing student progress
type StudentProgressUpdateRequest struct {
	ID                  int64   `json:"id" binding:"required" example:"1"`
	AudioUrl            string  `json:"audio_url" binding:"" example:"https://example.com/audio.mp3"`
	AiScore             float64 `json:"ai_score" binding:"" example:"85.5"`
	PronAccuracyScore   float64 `json:"pron_accuracy_score" binding:"" example:"80.0"`
	PronFluencyScore    float64 `json:"pron_fluency_score" binding:"" example:"90.0"`
	PronCompletionScore float64 `json:"pron_completion_score" binding:"" example:"85.0"`
	WordScores          string  `json:"word_scores" binding:"" example:"[{\"word\":\"hello\",\"score\":90}]"`
	AiSuggestions       string  `json:"ai_suggestions" binding:"" example:"Work on the pronunciation of 'th' sounds."`
	TeacherScore        float64 `json:"teacher_score" binding:"" example:"90.0"`
	TeacherFeedback     string  `json:"teacher_feedback" binding:"" example:"Good job, but work on your pronunciation."`
	TeacherId           string  `json:"teacher_id" binding:"" example:"teacher123"`
}

// TeacherFeedbackRequest - Request to submit teacher feedback
type TeacherFeedbackRequest struct {
	ID              int64   `json:"id" binding:"required" example:"1"`
	TeacherScore    float64 `json:"teacher_score" binding:"required" example:"90.0"`
	TeacherFeedback string  `json:"teacher_feedback" binding:"required" example:"Good job, but work on your pronunciation."`
	TeacherId       string  `json:"teacher_id" binding:"required" example:"teacher123"`
}

// StudentProgressDeleteRequest - Request to delete a student progress
type StudentProgressDeleteRequest struct {
	ID int64 `form:"id" binding:"required" example:"1"`
}
