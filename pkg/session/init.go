package session

import (
	"LarkBing/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var redisClient redis.UniversalClient

func getRedisLoginURL() string {
	loginInfo := config.C.Redis

	return fmt.Sprintf("%s:%d", loginInfo.Host, loginInfo.Port)
}

func ConnectRedis() {
	addr := make([]string, 0)
	addr = append(addr, getRedisLoginURL())
	redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    addr,
		DB:       config.C.Redis.DB,
		Password: config.C.Redis.Password,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	logrus.Info("Redis connected")
}
