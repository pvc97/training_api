package model

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

// Products - Mock list product
var Products = []Product{
	{ID: 1, Name: "Iphone 12", Price: 10000000, Quantity: 10},
	{ID: 2, Name: "Samsung Galaxy S21", Price: 12000000, Quantity: 5},
	{ID: 3, Name: "Xiaomi Redmi Note 10", Price: 5000000, Quantity: 20},
	{ID: 4, Name: "Oppo A94", Price: 7000000, Quantity: 15},
	{ID: 5, Name: "Vivo V21", Price: 8000000, Quantity: 25},
	{ID: 6, Name: "Realme 8", Price: 6000000, Quantity: 30},
	{ID: 7, Name: "Nokia 5.4", Price: 4000000, Quantity: 40},
	{ID: 8, Name: "Huawei P40", Price: 9000000, Quantity: 35},
	{ID: 9, Name: "Sony Xperia 1", Price: 11000000, Quantity: 45},
	{ID: 10, Name: "Google Pixel 5", Price: 13000000, Quantity: 50},
	{ID: 11, Name: "OnePlus 9", Price: 15000000, Quantity: 55},
	{ID: 12, Name: "Asus Zenfone 8", Price: 14000000, Quantity: 60},
	{ID: 13, Name: "Motorola Moto G100", Price: 16000000, Quantity: 65},
	{ID: 14, Name: "Lenovo Legion Duel 2", Price: 18000000, Quantity: 70},
	{ID: 15, Name: "BlackBerry Key2", Price: 20000000, Quantity: 75},
	{ID: 16, Name: "HTC U12+", Price: 22000000, Quantity: 80},
	{ID: 17, Name: "LG V60 ThinQ", Price: 24000000, Quantity: 85},
	{ID: 18, Name: "ZTE Axon 30 Ultra", Price: 26000000, Quantity: 90},
	{ID: 19, Name: "Meizu 18", Price: 28000000, Quantity: 95},
	{ID: 20, Name: "TCL 20 Pro 5G", Price: 30000000, Quantity: 100},
	{ID: 21, Name: "Alcatel 3L (2021)", Price: 32000000, Quantity: 105},
}
