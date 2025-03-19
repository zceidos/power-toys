package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 从环境变量获取 Redis 配置
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0, // 使用默认 DB
	})

	// 测试连接
	ctx := context.Background()
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}
	fmt.Println("Redis connected successfully")
}

// 设置带过期时间的缓存
func SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RDB.Set(ctx, key, json, expiration).Err()
}

// 获取缓存
func GetCache(ctx context.Context, key string, dest interface{}) error {
	val, err := RDB.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// 删除缓存
func DeleteCache(ctx context.Context, key string) error {
	return RDB.Del(ctx, key).Err()
}
