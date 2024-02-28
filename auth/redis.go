package main

import (
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
)

func ConnectRedis() (*redis.Client, error) {
	dbStr := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		// Handle the error if the conversion fails
		log.Fatalf("Error converting REDIS_DB to int: %v", err)
		return nil, err
	}

	log.Println(os.Getenv("REDIS_ADDR"))
	log.Println(os.Getenv("REDIS_PASSWORD"))

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	return RedisClient, nil
}
