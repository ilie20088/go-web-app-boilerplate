package utils

import "github.com/go-redis/redis"

var redisClient *redis.Client

// InitCache initializes redis client
func InitCache() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     GetCacheBDAddr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
