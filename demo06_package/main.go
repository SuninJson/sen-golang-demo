package main

//包的作用
//（1）区分相同名字的函数和变量等标识符
//（2）当程序文件很多时，可以很好的管理项目
//（3）控制函数、变量等访问范围

//可以使用 “package” 指令来打包

import "demo06_package/utils"

func main() {
	utils.HelloPackage()
}
