package controller

import (
	"blogSystem/bean"
	"blogSystem/initial"
	"blogSystem/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

type RegisterRequest struct {
	Username string `json:"userName" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 注册用户
func (controller *UserController) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDB bean.User
	if err := initial.DB.Where("username = ?", request.Username).First(&userDB).Error; err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username already exists"})
		return
	}
	if err := initial.DB.Where("email = ?", request.Email).First(&userDB).Error; err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email already exists"})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	request.Password = string(hashedPassword)

	user := bean.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := initial.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login 用户登录并认证
func (controller *UserController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser bean.User
	if err := initial.DB.Where("username = ?", request.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(storedUser.Model.ID, storedUser.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 响应token给客户端
	c.JSON(http.StatusOK, gin.H{"token": token})
}
