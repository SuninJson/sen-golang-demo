package main

import (
	"fmt"
	"sync"
)

var (
	mu1 sync.Mutex
	mu2 sync.Mutex
)

func worker() {
	mu1.Lock()
	defer mu1.Unlock()

	mu2.Lock()
	defer mu2.Unlock()

	fmt.Println("worker")
}

func main() {
	go worker()
	// 在主 goroutine 中获取锁的顺序与 worker() 不同，可能会导致死锁
	mu2.Lock()
	defer mu2.Unlock()

	mu1.Lock()
	defer mu1.Unlock()

	fmt.Println("main")
}
