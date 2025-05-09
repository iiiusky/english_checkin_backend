package v1

// Student data structure for API responses
type StudentDataItem struct {
	ID          uint   `json:"id"`
	OpenID      string `json:"open_id"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url"`
	PhoneNumber string `json:"phone_number"`
	ExpiryDate  int64  `json:"expiry_date"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updatedAt"`
	CreatedAt   string `json:"createdAt"`
}

// GetStudentRequest - Request to get a single student by ID
type GetStudentRequest struct {
	ID int64 `form:"id" binding:"required" example:"1"`
}

// GetStudentResponse - Response for a single student
type GetStudentResponse struct {
	Response
	Data StudentDataItem `json:"data"`
}

// GetStudentsRequest - Request for paginated student list
type GetStudentsRequest struct {
	Page        int    `form:"page" binding:"" example:"1"`
	PageSize    int    `form:"pageSize" binding:"required" example:"10"`
	Name        string `form:"name" binding:"" example:"张三"`
	PhoneNumber string `form:"phoneNumber" binding:"" example:"13800138000"`
	Status      string `form:"status" binding:"" example:"active"`
}

// GetStudentsResponseData - Data for paginated student list
type GetStudentsResponseData struct {
	List  []StudentDataItem `json:"list"`
	Total int64             `json:"total"`
}

// GetStudentsResponse - Response for paginated student list
type GetStudentsResponse struct {
	Response
	Data GetStudentsResponseData `json:"data"`
}

// StudentCreateRequest - Request to create a new student
type StudentCreateRequest struct {
	OpenID      string `json:"open_id" binding:"" example:"wx123456"`
	Name        string `json:"name" binding:"required" example:"张三"`
	AvatarURL   string `json:"avatar_url" binding:"" example:"https://example.com/avatar.jpg"`
	PhoneNumber string `json:"phone_number" binding:"" example:"13800138000"`
	ExpiryDate  int64  `json:"expiry_date" binding:"" example:"1714608000000"`
	Status      string `json:"status" binding:"required" example:"active"`
}

// StudentUpdateRequest - Request to update an existing student
type StudentUpdateRequest struct {
	ID          int64  `json:"id" binding:"required" example:"1"`
	Name        string `json:"name" binding:"required" example:"张三"`
	AvatarURL   string `json:"avatar_url" binding:"" example:"https://example.com/avatar.jpg"`
	PhoneNumber string `json:"phone_number" binding:"" example:"13800138000"`
	ExpiryDate  int64  `json:"expiry_date" binding:"" example:"1714608000000"`
	Status      string `json:"status" binding:"required" example:"active"`
}

// StudentDeleteRequest - Request to delete a student
type StudentDeleteRequest struct {
	ID int64 `form:"id" binding:"required" example:"1"`
}
