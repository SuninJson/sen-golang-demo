package main

import (
	"fmt"
	"sen-golang-demo/demo08_encapsulate/model"
)

func main() {
	jason := model.NewPerson("Jason")
	jason.SetAge(18)

	fmt.Println(jason)
}
