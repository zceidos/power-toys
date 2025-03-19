package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/skyeidos/power-toys/models"
)

// 定义刷新缓存的间隔时间
const cacheRefreshInterval = 30 * time.Second

// StartCacheRefresh 启动定时刷新缓存的 goroutine
func StartCacheRefresh() {
	go func() {
		for {
			// 等待指定的时间间隔
			time.Sleep(cacheRefreshInterval)

			// 刷新缓存逻辑
			refreshCache()
		}
	}()
}

// 刷新缓存的具体实现
func refreshCache() {
	ctx := context.Background()
	// 这里可以添加您需要刷新的缓存逻辑
	// 例如，重新从数据库加载用户权限并更新到 Redis
	var users []models.User
	if err := DB.Preload("Role.Permissions").Find(&users).Error; err != nil {
		log.Printf("Failed to refresh cache: %v", err)
		return
	}

	for _, user := range users {
		cacheKey := fmt.Sprintf("user_permissions:%d", user.ID)
		err := SetCache(ctx, cacheKey, user.Role.Permissions, cacheRefreshInterval)
		if err != nil {
			log.Printf("Failed to set cache for user %d: %v", user.ID, err)
		}
	}

	log.Println("Cache refreshed successfully")
}
