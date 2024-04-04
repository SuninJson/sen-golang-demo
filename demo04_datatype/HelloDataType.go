package main

import "fmt"

func main() {
	//数据类型可以从 小->大 也可以从 大->小
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n\n", i, n1, n2, n3)

	//被转换的是变量存储的数据，变量本身的数据类型不会发生变化
	fmt.Printf("i type is %T \n", i)
	fmt.Printf("n1 type is %T \n", n1)
	fmt.Printf("n2 type is %T \n", n2)
	fmt.Printf("n3 type is %T \n\n", n3)

	//大->小，编译时不会报错，但转换的结果会按溢出处理
	//例如 int64 -> int8
	var num1 int64 = 9999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)

	//uint是无符号位数，所以-1为2的64次方减一
	var num3 uint = 1
	var num4 uint = 2
	fmt.Printf("uint 1 - uint 2 = %v",num3-num4)
}
