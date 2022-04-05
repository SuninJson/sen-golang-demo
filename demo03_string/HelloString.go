package main

import "fmt"

func main() {
	// golang中字符串强制使用UTF-8编码
	// 字符串是不可变的
	var address string = "Hello String"
	fmt.Println(address)

	//字符串有两种表示形式
	//（1）双引号：会识别转义字符
	//（2）反引号：以字符串的原生形式输出，包括换行和特殊字符
	str1 := "双引号实例\n双引号实例"
	fmt.Println(str1)

	str2 := `go.mod

			module sen-golang-demo
			
			go 1.17
			`
	fmt.Println(str2)

	//字符串拼接
	//换行需要将 ‘+’ 留在上一行
	str3 := " demo01_hello" + " world1" +
		" demo01_hello" + " world2" +
		" demo01_hello" + " world3"
	fmt.Println(str3)

	//基本数据类型转String
	var str string

	var num1 int = 99
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q \n", str, str)

	var num2 float64 = 3.14
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q \n", str, str)

	var b bool = true
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q \n", str, str)
}
