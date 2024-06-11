package model

type CreateProductReq struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Cover    string `json:"cover"`
}
