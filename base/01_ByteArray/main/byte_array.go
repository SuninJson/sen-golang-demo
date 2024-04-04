package main

import "fmt"

func main() {
	s := "Hello, world!"
	b := []byte(s) // 将字符串转换为字节数组

	// 修改字节数组中的内容
	b[7] = 'W'
	b = append(b[:13], '!')

	// 将字节数组转换回字符串并输出
	fmt.Println(string(b))
}
