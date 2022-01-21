package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"testing"
	"time"
)

// 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
// Redis 基准测试 redis-benchmark 是一种实用工具，https://www.redis.com.cn/redis-benchmarks.html
// redis-benchmark [option] [option value]
// redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 10 -q
// 直接在服务器上执行上面命令，测试结果如下：
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 10 -q
// SET: 204081.64 requests per second
// GET: 200000.00 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 20 -q
// SET: 196078.44 requests per second
// GET: 192307.69 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 50 -q
// SET: 200000.00 requests per second
// GET: 188679.25 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 100 -q
// SET: 178571.42 requests per second
// GET: 185185.19 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 200 -q
// SET: 175438.59 requests per second
// GET: 188679.25 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 1000 -q
// SET: 123456.79 requests per second
// GET: 175438.59 requests per second
//
// [root@172-16-151-36 ~]# redis-benchmark -h 172.16.151.36 -p 6379 -t get,set -n 10000 -d 5000 -q
// SET: 86206.90 requests per second
// GET: 156250.00 requests per second

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

// 分别写入不同大小的，不同数量的数据，完成后分别使用redis-cli命令，info查看memory的信息
func TestRedis(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	set(10000, "1w", generateValue(10))
	set(10000, "1w", generateValue(100))
	set(10000, "1w", generateValue(1000))

	set(50000, "5w", generateValue(10))
	set(50000, "5w", generateValue(100))
	set(50000, "5w", generateValue(1000))

	set(100000, "10w", generateValue(10))
	set(100000, "10w", generateValue(100))
	set(100000, "10w", generateValue(1000))

	// 执行完成后，命令行观察 redis-cli 的 memory 信息
}

func set(num int, key string, value string) {
	for i := 0; i < num; i++ {
		err := redisDb.Set(ctx, key+"_"+strconv.Itoa(i), value, 600*time.Second).Err()
		if err != nil {
			log.Println(err)
		}
	}
}

func generateValue(size int) string {
	slice := make([]byte, size)
	for i := 0; i < size; i++ {
		slice[i] = '1'
	}
	return string(slice)
}
