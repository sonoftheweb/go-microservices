package handler

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewHandler(db *gorm.DB, redisClient *redis.Client) *Handler {
	return &Handler{
		DB:          db,
		RedisClient: redisClient,
	}
}
