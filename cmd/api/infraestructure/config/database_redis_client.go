package config

import (
	"os"
	"strconv"

	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/go-redis/redis"
)

const (
	RedisHost   = "REDIS_HOST"
	RedisSchema = "REDIS_SCHEMA"
	RedisExpire = "REDIS_EXPIRE"
)

func GetClient() *redis.Client {
	address := os.Getenv(RedisHost)
	schema, _ := strconv.Atoi(os.Getenv(RedisSchema))

	db := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       schema,
	})

	_, err := db.Ping().Result()
	if err != nil {
		logger.Error(err.Error(), err)
		_ = db.Close()
		panic("Database connection failed")
	}

	return db
}
