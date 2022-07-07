package test

import (
	"fmt"
	"testing"
	"time"
)

// 输出当前日期、时间戳
func TestTime1(t *testing.T) {
	timeobj := time.Now()
	timeStr := timeobj.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	fmt.Println(timeobj.Unix()) // 当前时间戳
	fmt.Println(timeobj.UnixMilli())
	fmt.Println(timeobj.UnixMicro())
	fmt.Println(timeobj.UnixNano())
}

// 时间戳转日期
func TestTime2(t *testing.T) {
	var unixTime int64 = 1644549697
	timeobj := time.Unix(unixTime, 0)
	timeStr := timeobj.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
}

// 日期转时间戳
func TestTime3(t *testing.T) {
	dateStr := "2018-09-18 04:14:56"
	tmp := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai") // 设置时区
	// loc := time.Local
	fmt.Println(loc)
	timeobj, _ := time.ParseInLocation(tmp, dateStr, loc) // 2006-01-02 15:04:05是转换的格式，如php的"Y-m-d H:i:s"
	fmt.Println(timeobj)
	fmt.Println(timeobj.Unix())
}

func TestTime33(t *testing.T) {
	dateStr := "2018-09-18 04:14:56"
	tmp := "2006-01-02 15:04:05"
	timeobj, _ := time.Parse(tmp, dateStr)
	fmt.Println(timeobj)
}

func TestTime333(t *testing.T) {
	dateStr := "2022-02-17 16:14:56"
	tmp := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// loc := time.Local
	fmt.Println(loc)
	timeobj, _ := time.ParseInLocation(tmp, dateStr, loc)
	fmt.Println(timeobj)
	fmt.Println(time.Now().Sub(timeobj))

	ttt := time.Now()
	fmt.Println(ttt.Unix() - timeobj.Unix())
}

// 时间单位
func TestTime4(t *testing.T) {
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
}

// 几天后
func TestTime5(t *testing.T) {
	timeobj := time.Now()
	fmt.Println(timeobj)
	timeobj2 := timeobj.Add(time.Hour * 3)
	timeobj3 := timeobj.AddDate(0, 0, 3)
	fmt.Println(timeobj2)
	fmt.Println(timeobj3)
}

// 定时器
func TestTime6(t *testing.T) {
	// 生成一个定时器，每隔3秒去执行一个任务
	ticker := time.NewTicker(time.Second * 2)
	count := 0

	//for {
	//	if data, ok := <-ticker.C; ok {
	//		count++
	//		fmt.Println(data)
	//		if count > 5 {
	//			ticker.Stop()
	//			break
	//		}
	//	} else {
	//		break
	//	}
	//}

	for val := range ticker.C {
		fmt.Println(val)
		count++
		if count > 5 {
			ticker.Stop()
			break
		}
	}
}
