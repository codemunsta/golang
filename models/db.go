package models

import "github.com/go-redis/redis"

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
