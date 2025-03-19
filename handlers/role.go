package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyeidos/power-toys/config"
	"github.com/skyeidos/power-toys/models"
)

type CreateRoleRequest struct {
	Name          string `json:"name" binding:"required"`
	PermissionIDs []uint `json:"permission_ids"`
}

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{
		Name: req.Name,
	}

	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	// 如果提供了权限ID，则关联权限
	if len(req.PermissionIDs) > 0 {
		var permissions []models.Permission
		if err := config.DB.Find(&permissions, req.PermissionIDs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission IDs"})
			return
		}
		config.DB.Model(&role).Association("Permissions").Replace(&permissions)
	}

	c.JSON(http.StatusCreated, role)
}

// GetRoles 获取角色列表
func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := config.DB.First(&role, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	role.Name = req.Name
	if err := config.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	// 更新权限
	if len(req.PermissionIDs) > 0 {
		var permissions []models.Permission
		if err := config.DB.Find(&permissions, req.PermissionIDs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission IDs"})
			return
		}
		config.DB.Model(&role).Association("Permissions").Replace(&permissions)
	}

	// 清除相关用户的权限缓存
	var users []models.User
	config.DB.Where("role_id = ?", role.ID).Find(&users)
	for _, user := range users {
		ClearUserPermissionsCache(user.ID)
	}

	c.JSON(http.StatusOK, role)
}
