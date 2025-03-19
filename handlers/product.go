package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyeidos/power-toys/config"
	"github.com/skyeidos/power-toys/models"
)

// @Summary     创建商品
// @Description 创建新商品
// @Tags        商品管理
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "Bearer 用户令牌"
// @Param       request body models.Product true "商品信息"
// @Success     201 {object} models.Product "创建成功的商品信息"
// @Failure     400 {object} map[string]string "请求参数错误"
// @Failure     401 {object} map[string]string "未授权"
// @Failure     403 {object} map[string]string "权限不足"
// @Router      /products [post]
// @Security    Bearer
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// @Summary     获取商品列表
// @Description 获取所有商品
// @Tags        商品管理
// @Produce     json
// @Param       Authorization header string true "Bearer 用户令牌"
// @Success     200 {array} models.Product "商品列表"
// @Failure     401 {object} map[string]string "未授权"
// @Failure     403 {object} map[string]string "权限不足"
// @Router      /products [get]
// @Security    Bearer
func GetProducts(c *gin.Context) {
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary     获取商品详情
// @Description 根据ID获取商品详情
// @Tags        商品管理
// @Produce     json
// @Param       Authorization header string true "Bearer 用户令牌"
// @Param       id path int true "商品ID"
// @Success     200 {object} models.Product "商品详情"
// @Failure     400 {object} map[string]string "ID格式错误"
// @Failure     401 {object} map[string]string "未授权"
// @Failure     403 {object} map[string]string "权限不足"
// @Failure     404 {object} map[string]string "商品不存在"
// @Router      /products/{id} [get]
// @Security    Bearer
func GetProduct(c *gin.Context) {
	var product models.Product
	result := config.DB.First(&product, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct 更新商品
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct 删除商品
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
