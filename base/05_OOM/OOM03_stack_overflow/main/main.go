package main

import "fmt"

// 正确设置递归基准条件
func main() {
	fmt.Println(fibonacci(5))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	// 设置递归基准条件
	return fibonacci(n-1) + fibonacci(n-2)
}
