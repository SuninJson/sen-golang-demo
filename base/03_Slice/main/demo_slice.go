package main

import "fmt"

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	// 定义整数数组
	arr := []int{1, 2, 3, 4, 5}
	newArr := arr[1:3]

	fmt.Printf("newArr: %v, len=%d, cap=%d\n", newArr, len(newArr), cap(newArr))

	// 调用求和函数
	s := sum(newArr)

	fmt.Println("sum:", s)
}
