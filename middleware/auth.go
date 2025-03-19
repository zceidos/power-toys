package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skyeidos/power-toys/config"
	"github.com/skyeidos/power-toys/models"
	"github.com/skyeidos/power-toys/utils"
)

type CachedPermissions struct {
	UserID      uint
	Permissions []models.Permission
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 尝试从 Redis 获取权限信息
		var cachedPerms CachedPermissions
		cacheKey := fmt.Sprintf("user_permissions:%d", claims.UserID)
		ctx := context.Background()

		err = config.GetCache(ctx, cacheKey, &cachedPerms)
		if err != nil {
			// 缓存未命中，从数据库获取
			var user models.User
			if err := config.DB.Preload("Role.Permissions").First(&user, claims.UserID).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				c.Abort()
				return
			}

			// 更新缓存
			cachedPerms = CachedPermissions{
				UserID:      user.ID,
				Permissions: user.Role.Permissions,
			}
			err = config.SetCache(ctx, cacheKey, cachedPerms, 30*time.Minute)
			if err != nil {
				fmt.Printf("Failed to set cache: %v\n", err)
			}
		}

		// 检查权限
		path := c.Request.URL.Path
		method := c.Request.Method

		hasPermission := false
		for _, perm := range cachedPerms.Permissions {
			if perm.Path == path && perm.Method == method {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
