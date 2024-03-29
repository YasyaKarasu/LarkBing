package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func SetSession(key string, value any) error {
	err := redisClient.Set(context.Background(), key, value, time.Hour*24*7).Err()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func GetSessionString(key string) string {
	content, err := redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		logrus.WithField("ID", key).Warn("key does not exist")
		return ""
	} else if err != nil {
		logrus.WithField("ID", key).Error(err)
		return ""
	} else {
		return content
	}
}

func GetSessionBool(key string) (bool, bool) {
	content, err := redisClient.Get(context.Background(), key).Bool()
	if err == redis.Nil {
		logrus.WithField("ID", key).Warn("key does not exist")
		return false, false
	} else if err != nil {
		logrus.WithField("ID", key).Error(err)
		return false, false
	} else {
		return content, true
	}
}

func ClearSession(key string) {
	err := redisClient.Del(context.Background(), key).Err()
	if err != nil {
		logrus.Error(err)
	}
}
