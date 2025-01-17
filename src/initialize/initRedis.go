package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"userManageSystem-blog/src/pkg/globals"
)

func InitRedis(c *globals.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Config.Host,
		Password: "",
		DB:       c.Redis.Config.Db,
		PoolSize: c.Redis.Config.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := rdb.Ping(ctx).Result() // 检查连接redis是否成功
	if err != nil {
		fmt.Printf("Connect HttpFail: %v \n\n", err)
		panic(err)
	}
	fmt.Println("redis连接成功")
	return nil
}
