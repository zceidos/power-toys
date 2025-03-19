package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyeidos/power-toys/config"
	"github.com/skyeidos/power-toys/models"
	"github.com/skyeidos/power-toys/utils"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"admin123"`
}

// @Summary     用户登录
// @Description 用户登录并获取token
// @Tags        认证
// @Accept      json
// @Produce     json
// @Param       request body LoginRequest true "登录信息"
// @Success     200 {object} map[string]interface{} "token和用户信息"
// @Failure     400 {object} map[string]string "错误信息"
// @Failure     401 {object} map[string]string "认证失败"
// @Router      /login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 生成 token
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

// 在更新用户角色或权限时调用此函数
func ClearUserPermissionsCache(userID uint) error {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("user_permissions:%d", userID)
	return config.DeleteCache(ctx, cacheKey)
}
