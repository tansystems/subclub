package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Модуль работы с Redis: хранение сессий, подписок, пользователей.

// Сохраняет подписку пользователя
func SaveSubscription(redisClient *redis.Client, userID string) error {
	ctx := context.Background()
	return redisClient.Set(ctx, "subscribed:"+userID, "1", 0).Err()
}

// Проверяет наличие подписки
func CheckSubscription(redisClient *redis.Client, userID string) (bool, error) {
	ctx := context.Background()
	val, err := redisClient.Get(ctx, "subscribed:"+userID).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val == "1", nil
}

// Сохраняет пользователя
func SaveUser(redisClient *redis.Client, username, password string) error {
	ctx := context.Background()
	return redisClient.Set(ctx, "user:"+username, password, 0).Err()
}

// Получает пароль пользователя
func GetUser(redisClient *redis.Client, username string) (string, error) {
	ctx := context.Background()
	return redisClient.Get(ctx, "user:"+username).Result()
}
