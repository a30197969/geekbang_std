package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var redisDb *redis.Client
var ctx context.Context

// 根据Redis配置初始化一个客户端
func initRedisClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "172.16.151.36:6379", // redis地址
		Password: "",                   // redis密码，没有则留空
		DB:       0,                    // 默认数据库，默认是0
	})
	ctx = context.Background()
	_, err = redisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
func TestName(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	err = redisDb.Set(ctx, "go-test-1", "大家都很好", 600*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}
