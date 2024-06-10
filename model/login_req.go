package model

type LoginRequest struct {
	TaxCode  int    `json:"tax_code"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
