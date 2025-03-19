package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skyeidos/power-toys/config"
	_ "github.com/skyeidos/power-toys/docs"
	"github.com/skyeidos/power-toys/handlers"
	"github.com/skyeidos/power-toys/metrics"
	"github.com/skyeidos/power-toys/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Power Toys API
// @version         1.0
// @description     Power Toys 后台管理系统 API 文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// 初始化数据库和 Redis
	config.InitDB()
	config.InitRedis()

	// 启动定时刷新缓存
	config.StartCacheRefresh()

	r := gin.Default()

	// 添加 Prometheus 中间件
	r.Use(metrics.PrometheusMiddleware())

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus 指标接口
	r.GET("/metrics", handlers.PrometheusHandler())

	// 公开路由
	r.POST("/login", handlers.Login)

	// 需要认证的路由
	products := r.Group("/products")
	products.Use(middleware.AuthMiddleware())
	{
		products.POST("", handlers.CreateProduct)
		products.GET("", handlers.GetProducts)
		products.GET("/:id", handlers.GetProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	// 用户管理路由
	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.POST("", handlers.CreateUser)
		users.GET("", handlers.GetUsers)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	// 角色管理路由
	roles := r.Group("/roles")
	roles.Use(middleware.AuthMiddleware())
	{
		roles.POST("", handlers.CreateRole)
		roles.GET("", handlers.GetRoles)
		roles.PUT("/:id", handlers.UpdateRole)
	}

	r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
