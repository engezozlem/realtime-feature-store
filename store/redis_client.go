package store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := Rdb.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Redis bağlantısı kurulamadı: %v", err)
	}
}
