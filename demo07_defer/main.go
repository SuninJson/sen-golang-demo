package main

import "fmt"

func sum(n1 int, n2 int) int {
	//当识别到defer时，会暂不执行，会将defer所标识的语句压到defer栈中，
	//在函数执行完成时，再从defer栈按先入后出的方式执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok1 n2 = ", n2)

	result := n1 + n2
	fmt.Println("result =", result)

	return result
}

func main() {
	result := sum(10, 20)
	fmt.Println(result)
}
