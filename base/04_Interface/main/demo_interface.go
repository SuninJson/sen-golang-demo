package main

import (
	"fmt"
	"testing"
)

// Animal 定义一个接口
type Animal interface {
	Speak() string
}

// Dog 定义一个结构体
type Dog struct{}

// Speak 实现接口方法
func (d Dog) Speak() string {
	return "汪汪!"
}

// Cat 定义另一个结构体
type Cat struct{}

// Speak 实现接口方法
func (c Cat) Speak() string {
	return "喵喵!"
}

func TestAnimal(t *testing.T) {
	// 创建一个切片，其中包含了 Dog 和 Cat 两种类型
	animals := []Animal{Dog{}, Cat{}}

	// 循环遍历所有的动物并测试它们的 Speak() 方法是否正确
	for _, animal := range animals {
		if animal.Speak() == "" {
			t.Errorf("Speak()方法返回值为空")
		}
	}
}

func main() {
	// 将狗和猫实例化，并且存储在 Animal 接口变量中
	var animal Animal

	animal = Dog{}
	fmt.Println(animal.Speak())

	animal = Cat{}
	fmt.Println(animal.Speak())
}
