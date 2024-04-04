package main

import (
	"fmt"
	"sync"
)

func foo() {
	for i := 0; i < 1000000; i++ {
		// 使用 sync.Once 确保只注册一次
		once := sync.Once{}
		once.Do(func() {
			defer func() {
				fmt.Println("defer")
			}()
		})
	}
}
