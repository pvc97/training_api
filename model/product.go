package model

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Cover    string `json:"cover"`
}

// Products - Mock list product
var Products = []Product{
	{ID: 1, Name: "Iphone 12", Price: 10000000, Quantity: 10, Cover: "https://cdn.tgdd.vn/Products/Images/42/326016/vivo-y28-vang-thumb-600x600.jpg"},
	{ID: 2, Name: "Samsung Galaxy S21", Price: 12000000, Quantity: 5, Cover: "https://cdn.tgdd.vn/Products/Images/42/305658/iphone-15-pro-max-blue-thumbnew-600x600.jpg"},
	{ID: 3, Name: "Xiaomi Redmi Note 10", Price: 5000000, Quantity: 20, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 4, Name: "Oppo A94", Price: 7000000, Quantity: 15, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 5, Name: "Vivo V21", Price: 8000000, Quantity: 25, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 6, Name: "Realme 8", Price: 6000000, Quantity: 30, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 7, Name: "Nokia 5.4", Price: 4000000, Quantity: 40, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 8, Name: "Huawei P40", Price: 9000000, Quantity: 35, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 9, Name: "Sony Xperia 1", Price: 11000000, Quantity: 45, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 10, Name: "Google Pixel 5", Price: 13000000, Quantity: 50, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 11, Name: "OnePlus 9", Price: 15000000, Quantity: 55, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 12, Name: "Asus Zenfone 8", Price: 14000000, Quantity: 60, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 13, Name: "Motorola Moto G100", Price: 16000000, Quantity: 65, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 14, Name: "Lenovo Legion Duel 2", Price: 18000000, Quantity: 70, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 15, Name: "BlackBerry Key2", Price: 20000000, Quantity: 75, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 16, Name: "HTC U12+", Price: 22000000, Quantity: 80, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 17, Name: "LG V60 ThinQ", Price: 24000000, Quantity: 85, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 18, Name: "ZTE Axon 30 Ultra", Price: 26000000, Quantity: 90, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 19, Name: "Meizu 18", Price: 28000000, Quantity: 95, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 20, Name: "TCL 20 Pro 5G", Price: 30000000, Quantity: 100, Cover: "https://cdn.tgdd.vn/Products/Images/42/249948/samsung-galaxy-s23-ultra-green-thumbnew-600x600.jpg"},
	{ID: 21, Name: "Alcatel 3L (2021)", Price: 32000000, Quantity: 105, Cover: "https://cdn.tgdd.vn/Products/Images/42/322096/samsung-galaxy-a55-5g-xanh-thumb-1-600x600.jpg"},
}
