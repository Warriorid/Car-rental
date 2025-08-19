package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


func NewClient() (*redis.Client, error) {
	
	password := os.Getenv("REDIS_PASSWORD")
	addr := fmt.Sprintf("redis:%s", viper.GetString("redis.port"))
	opts := &redis.Options{
		Addr: addr,
		Password: password,
		DB: 0,
	}
	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logrus.Fatalf("Ping error: %s", err)
		return nil, err
	}
	return client, nil

}


