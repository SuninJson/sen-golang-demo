// 每个go文件需要属于一个包
package main

// 用到的包需要进行导入
// 如果用到的包没有用到，代码无法编译通过
import "fmt"

// main函数来执行程序
func main() {
	// golang结尾没有分号，一行只能写一条语句
	fmt.Println("demo01_hello world!")
	fmt.Println("demo01_hello", //可以使用逗号换行
		"world!")
}
