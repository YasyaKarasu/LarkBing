package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func SetContextSession(ID string, content string) error {
	err := redisClient.Set(context.Background(), ID, content, time.Hour*24*7).Err()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func GetContextSession(ID string) string {
	content, err := redisClient.Get(context.Background(), ID).Result()
	if err == redis.Nil {
		logrus.WithField("ID", ID).Warn("key does not exist")
		return ""
	} else if err != nil {
		logrus.WithField("ID", ID).Error(err)
		return ""
	} else {
		return content
	}
}

func ClearContextSession(ID string) {
	err := redisClient.Del(context.Background(), ID).Err()
	if err != nil {
		logrus.Error(err)
	}
}
