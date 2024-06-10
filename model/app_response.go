package model

type appResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewAppResponse(success bool, message string, data any) *appResponse {
	return &appResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}
