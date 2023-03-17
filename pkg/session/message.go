package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func StoreMessageID(messageID string) error {
	err := redisClient.Set(context.Background(), messageID, true, time.Hour*24*7).Err()
	if err != nil {
		logrus.WithField("messageID", messageID).Error(err)
		return err
	}
	return nil
}

func QueryMessageID(messageID string) bool {
	_, err := redisClient.Get(context.Background(), messageID).Bool()
	if err == redis.Nil {
		return false
	} else if err != nil {
		logrus.WithField("messageID", messageID).Error(err)
		return false
	} else {
		return true
	}
}
