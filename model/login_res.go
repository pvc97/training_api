package model

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
