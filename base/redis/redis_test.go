package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"math"
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

// 字符串练习
func TestName(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var key = "go-test-string-1"
	err = redisDb.Set(ctx, key, "大家都很好1", 600*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
	value := redisDb.Get(ctx, key)
	fmt.Printf("%#v\n", value)
}

// 写入集合
func TestName2(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	l1 := redis.Z{
		Score:  91.96,
		Member: "Object-C",
	}
	num, err := redisDb.ZAdd(ctx, zsetKeyName, &l1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	//languages := []redis.Z{
	//	redis.Z{Score: 90.0, Member: "Golang"},
	//	redis.Z{Score: 92.78, Member: "Java"},
	//	redis.Z{Score: 91.96, Member: "PHP"},
	//	redis.Z{Score: 94.08, Member: "C"},
	//	redis.Z{Score: 93.22, Member: "JS"},
	//}
	//for _, language := range languages {
	//	num, err := redisDb.ZAdd(ctx, zsetKeyName, &language).Result()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(num)
	//}
}

// 集合 增加元素分值
func TestName2_1(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	num, err := redisDb.ZIncrBy(ctx, zsetKeyName, -1, "javaScript").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", num)
}

// 集合 分值排序：获取升序和降序的列表数据
func TestName2_2(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	// 降序
	resRev, err := redisDb.ZRevRange(ctx, zsetKeyName, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("降序 ZRevRange：%#v\n", resRev)
	// 升序
	res, err := redisDb.ZRange(ctx, zsetKeyName, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("升序 ZRange%#v\n", res)
}

// 集合 分值排序：根据score分值范围，获取升序和降序的列表数据
func TestName2_3(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	// 降序
	resRev, err := redisDb.ZRevRangeByScore(ctx, zsetKeyName, &redis.ZRangeBy{
		Min: "91",
		Max: "92",
	}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("降序 ZRevRange：%#v\n", resRev)
	// 升序
	res, err := redisDb.ZRangeByScore(ctx, zsetKeyName, &redis.ZRangeBy{
		Min: "90",
		Max: "91",
	}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("升序 ZRange%#v\n", res)
}

// 获取元素总个数
func TestName2_4(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	num, err := redisDb.ZCard(ctx, zsetKeyName).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("元素总个数：%#v\n", num)
}

// 获取区间内元素个数
func TestName2_5(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	num, err := redisDb.ZCount(ctx, zsetKeyName, "90", "92").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("元素个数：%#v\n", num)
}

// 获取某个元素的score
func TestName2_6(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	score, err := redisDb.ZScore(ctx, zsetKeyName, "Object-C").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("分数：%#v\n", score)
}

// 获取某个元素在集合中的排名
func TestName2_7(t *testing.T) {
	err := initRedisClient()
	if err != nil {
		log.Fatalln(err)
	}
	var zsetKeyName = "go-test-zset-1"
	rank, err := redisDb.ZRevRank(ctx, zsetKeyName, "Object-C").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("排名：%#v\n", rank)
	fmt.Println(math.Pow(10, 14))
}
