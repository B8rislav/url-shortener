package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 6 * time.Hour

func InitStore(ctx context.Context) error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	fmt.Printf("\nRedis started successfully, message: %v", pong)
	storeService.redisClient = redisClient
	return nil
}

func SaveUrl(ctx context.Context, shortUrl string, originalUrl string, userId string) error {
	log.Println(shortUrl, originalUrl, userId)
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
