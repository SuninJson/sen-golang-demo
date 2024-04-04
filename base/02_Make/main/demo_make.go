package main

import "fmt"

func main() {
	// 生成切片
	s := make([]int, 3, 5)
	fmt.Printf("slice: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// 生成映射
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	fmt.Println("map:", m)

	// 生成通道
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for i := range ch {
		fmt.Println("channel value:", i)
	}

	// 生成数组
	a := make([]int, 3)
	b := [3]int{1, 2, 3}
	copy(a, b[:])
	fmt.Printf("array: %v\n", a)

	// 生成二维切片
	matrix := make([][]int, 2)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 1
	matrix[1][2] = 2
	fmt.Println("matrix:", matrix)
}
