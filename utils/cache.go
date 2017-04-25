package utils

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// RedisClient is a global variable used to work with Redis
var RedisClient *redis.Client

// InitCache initializes redis client
func InitCache() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetCacheBDAddr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := PingRedis()
	if err != nil {
		Logger.Error("[Cache]", zap.Error(err))
	}
}

// PingRedis checks if Redis is available
var PingRedis = func() error {
	statusCmd := RedisClient.Ping()
	return statusCmd.Err()
}
