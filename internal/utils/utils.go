package utils

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[1;33m"
)

// ErrorResponse struct to respond back error message
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse struct to respond back success message
type SuccessResponse struct {
	Success string `json:"success"`
}

// HealthResponse struct to respond back to health request
type HealthResponse struct {
	ActiveTasks    uint `json:"active_tasks"`
	MaxActiveTasks uint `json:"max_active_tasks"`
}
