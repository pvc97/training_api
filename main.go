package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"training_api/model"
)

const fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != fakeToken {
			c.JSON(http.StatusUnauthorized, model.NewAppResponse(false, "Token không hợp lệ", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Ping
	r.GET("/reset", func(c *gin.Context) {
		model.Products = append([]model.Product{}, model.InitProducts...)
		c.JSON(http.StatusOK, gin.H{
			"message": "Reset data success",
		})
	})

	// Login
	r.POST("/login", login)

	// Login2
	r.POST("/login2", login2)

	// Get products
	r.GET("/products", AuthMiddleware(), getProducts)

	// Get product detail
	r.GET("/products/:id", AuthMiddleware(), getProductDetail)

	// Create product
	r.POST("/products", AuthMiddleware(), createProduct)

	// Delete product
	r.DELETE("/products/:id", AuthMiddleware(), deleteProduct)

	// Update product
	r.PUT("/products/:id", AuthMiddleware(), updateProduct)

	_ = r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
		c.JSON(http.StatusOK, model.NewLoginSuccessResponse(true, "Đăng nhập thành công", fakeToken))
		return
	} else {
		c.JSON(http.StatusUnauthorized, model.NewAppResponse(false, "Tên đăng nhập hoặc mật khẩu không đúng", nil))
		return
	}
}

func login2(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.NewAppResponse(false, "Dữ liệu không hợp lệ", nil))
		return
	}
	if loginRequest.TaxCode == 1111111111 &&
		loginRequest.UserName == "demo" &&
		loginRequest.Password == "123456" {
		c.JSON(http.StatusOK, model.NewAppResponse(true, "Đăng nhập thành công", model.LoginSuccessResponse2{
			Token: fakeToken,
		}))
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

	largestProductId := 0
	for _, product := range model.Products {
		if product.ID >= largestProductId {
			largestProductId = product.ID
		}
	}

	newId := largestProductId + 1
	newProduct := model.Product{
		ID:       newId,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Cover:    product.Cover,
	}

	// Add new product to the beginning of the list
	model.Products = append([]model.Product{newProduct}, model.Products...)

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
