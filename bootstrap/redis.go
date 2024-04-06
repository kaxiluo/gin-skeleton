package bootstrap

import (
	"context"
	"fmt"
	"gin-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:             global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password:         global.App.Config.Redis.Password, // no password set
		DB:               global.App.Config.Redis.DB,       // use default DB
		DisableIndentity: true,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis connect ping failed")
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
