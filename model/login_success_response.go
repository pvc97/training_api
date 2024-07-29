package model

type LoginSuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewLoginSuccessResponse(success bool, message, token string) *LoginSuccessResponse {
	return &LoginSuccessResponse{
		Success: success,
		Message: message,
		Token:   token,
	}
}
