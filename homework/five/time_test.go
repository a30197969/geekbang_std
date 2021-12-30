package five

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	var waitFiveHundredMilliseconds int64 = 500
	startingTime := time.Now()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now()
	var duration time.Duration = endingTime.Sub(startingTime)
	fmt.Printf("%+v\n", startingTime)
	fmt.Printf("%+v\n", endingTime)
	fmt.Printf("%+v\n", duration)
	var durationAsInt64 = int64(duration) // 10毫秒的Durantion类型对象转换为int64类型时，实际上得到的是10,000,000
	if durationAsInt64 >= waitFiveHundredMilliseconds {
		fmt.Printf("Time Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMilliseconds, durationAsInt64)
	} else {
		fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMilliseconds, durationAsInt64)
	}
}
func TestDuratio2(t *testing.T) {
	var duration_Milliseconds time.Duration = 500 * time.Millisecond
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute
	fmt.Printf("Milli [%v]\nSeconds [%v]\nMinute [%v]\n", duration_Milliseconds, duration_Seconds, duration_Minute)
}
func TestDuratio3(t *testing.T) {
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute
	var float64_Seconds float64 = duration_Seconds.Seconds()
	var float64_Minutes float64 = duration_Minute.Minutes()
	fmt.Printf("Seconds [%.3f]\nMinutes [%.2f]\n", float64_Seconds, float64_Minutes)
}
func TestDuratio4(t *testing.T) {
	var waitFiveHundredMilliseconds time.Duration = 500 * time.Millisecond
	startingTime := time.Now().UTC()
	time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()
	var duration time.Duration = endingTime.Sub(startingTime)
	if duration >= waitFiveHundredMilliseconds {
		fmt.Printf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", waitFiveHundredMilliseconds, duration, duration.Nanoseconds()/1e6, duration.Seconds())
	}
}
