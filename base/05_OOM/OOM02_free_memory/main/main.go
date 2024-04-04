package main

import (
	"fmt"
	"runtime"
)

type User struct {
	ID   int
	Name string
}

// 正确释放内存
func main() {
	u := &User{ID: 1, Name: "John"}
	// 注册终结器函数，在对象不再使用时释放内存
	runtime.SetFinalizer(u, func(u *User) {
		fmt.Printf("Free memory for object with ID %d\n", u.ID)
	})
	// 使用对象
	fmt.Println(u.Name)
}
