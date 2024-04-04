package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// OOM场景一：没有及时关闭文件句柄
func main() {
	// 获取当前文件所在目录
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println("Current directory:", dir)

	// 读取 oom.md 文件
	f, err := os.Open(filepath.Join(dir, "05_OOM", "oom.md"))
	if err != nil {
		log.Fatal(err)
	}
	// 在函数结束时延迟关闭文件句柄，若没有及时关闭文件句柄，可能会导致内存溢出
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// 读取文件内容
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
