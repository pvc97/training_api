package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"training_api/model"
)

func main() {
	r := gin.Default()
	r.POST("/login", login)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.LoginResponse{Success: false, Message: "Dữ liệu không hợp lệ"})
		return
	}
	if loginRequest.TaxCode == 1111111111 &&
		loginRequest.UserName == "demo" &&
		loginRequest.Password == "123456" {
		c.JSON(http.StatusOK, model.LoginResponse{Success: true, Message: "Đăng nhập thành công"})
	} else {
		c.JSON(http.StatusUnauthorized, model.LoginResponse{Success: false, Message: "Đăng nhập thất bại"})
	}
}
