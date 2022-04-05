package main

import "fmt"

//<a href = "https://www.processon.com/diagraming/623f70d6e0b34d0730dc05ae">图示</a>
func main() {
	var i int = 10
	fmt.Println("i的值：", i)
	fmt.Println("i的地址：", &i)

	//ptr是指针变量，类型是一个指向int的指针，本身的值是&i
	var ptr *int = &i
	fmt.Println("ptr的值：", ptr)
	fmt.Println("ptr的地址：", &ptr)
	fmt.Println("ptr的指向的地址的值：", *ptr)

	//每一个值类型都有对应的指针的类型，例如int对应*int，float对应*float
	//常见的值类型有int,float,bool,string,array和struct
	//值类型的变量直接存储值，内存通常在栈中分配
	//常见的引用类型有指针、slice、map、管道、interface
	//引用类型的变量存储的是一个内存地址，这个地址对应的空间才真正存储数据，
	//引用类型的变量的内存通常在堆上分配，当没有任何变量引用，该地址对应的数据空间就会成为垃圾，由GC回收

}
