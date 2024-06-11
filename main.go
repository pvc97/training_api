package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"training_api/model"
)

func main() {
	r := gin.Default()

	// Ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Login
	r.POST("/login", login)

	// Get products
	r.GET("/products", getProducts)

	// Get product detail
	r.GET("/products/:id", getProductDetail)

	// Create product
	r.POST("/products", createProduct)

	// Delete product
	r.DELETE("/products/:id", deleteProduct)

	// Update product
	r.PUT("/products/:id", updateProduct)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}
	if loginRequest.TaxCode == 1111111111 &&
		loginRequest.UserName == "demo" &&
		loginRequest.Password == "123456" {
		c.JSON(http.StatusOK, model.NewAppResponse(true, "Đăng nhập thành công", nil))
		return
	} else {
		c.JSON(http.StatusUnauthorized, model.NewAppResponse(false, "Tên đăng nhập hoặc mật khẩu không đúng", nil))
		return
	}
}

func getProducts(c *gin.Context) {
	var pagingData model.Paging
	if err := c.ShouldBind(&pagingData); err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}
	pagingData.Fill()

	var products []model.Product

	productLen := len(model.Products)
	start := pagingData.Size * (pagingData.Page - 1)
	end := start + pagingData.Size

	if start > productLen {
		start = productLen
	}

	if end > productLen {
		end = productLen
	}

	products = model.Products[start:end]
	c.JSON(http.StatusOK, model.NewAppResponse(true, "Lấy danh sách sản phẩm thành công", products))
}

func getProductDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}
	for _, product := range model.Products {
		if id == product.ID {
			c.JSON(http.StatusOK, model.NewAppResponse(true, "Lấy thông tin sản phẩm thành công", product))
			return
		}
	}
	c.JSON(http.StatusNotFound, model.NewAppResponse(false, "Không tìm thấy sản phẩm", nil))
}

func createProduct(c *gin.Context) {
	var product model.CreateProductReq
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}

	newId := 1
	if len(model.Products) > 0 {
		newId = model.Products[len(model.Products)-1].ID + 1
	}
	newProduct := model.Product{
		ID:       newId,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	model.Products = append(model.Products, newProduct)

	c.JSON(http.StatusCreated, model.NewAppResponse(true, "Tạo sản phẩm thành công", newProduct))
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}

	for i, product := range model.Products {
		if id == product.ID {
			model.Products = append(model.Products[:i], model.Products[i+1:]...)
			c.JSON(http.StatusOK, model.NewAppResponse(true, "Xóa sản phẩm thành công", nil))
			return
		}
	}
	c.JSON(http.StatusNotFound, model.NewAppResponse(false, "Không tìm thấy sản phẩm", nil))
}

func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}

	var product model.CreateProductReq
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}

	for i, p := range model.Products {
		if id == p.ID {
			model.Products[i].Name = product.Name
			model.Products[i].Price = product.Price
			model.Products[i].Quantity = product.Quantity
			model.Products[i].Cover = product.Cover
			c.JSON(http.StatusOK, model.NewAppResponse(true, "Cập nhật sản phẩm thành công", model.Products[i]))
			return
		}
	}
	c.JSON(http.StatusNotFound, model.NewAppResponse(false, "Không tìm thấy sản phẩm", nil))
}
