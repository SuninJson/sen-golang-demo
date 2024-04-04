package main

import (
	"sync"
	"time"
)

func main() {

}

type CounterLimiter struct {
	Interval int64
	LastTime time.Time
	MaxCount int
	lock     sync.Mutex
	ReqCount int
}

func NewLimiter(interval int64, maxCount int) *CounterLimiter {
	return &CounterLimiter{
		//在Interval所表示的时间段内进行计数
		Interval: interval,
		LastTime: time.Now(),
		MaxCount: maxCount,
		lock:     sync.Mutex{},
		ReqCount: 0,
	}
}

// Check 计数器限流
func (limiter *CounterLimiter) Check() bool {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()
	now := time.Now()
	if now.Unix()-limiter.LastTime.Unix() > limiter.Interval {
		limiter.LastTime = now
		limiter.ReqCount = 0
	}
	if limiter.ReqCount < limiter.MaxCount {
		limiter.ReqCount += 1
		return true
	}
	return false
}
