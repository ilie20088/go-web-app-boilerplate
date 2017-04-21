package utils

import "github.com/go-redis/redis"

// RedisClient is a global variable used to work with Redis
var RedisClient *redis.Client

// InitCache initializes redis client
func InitCache() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetCacheBDAddr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
