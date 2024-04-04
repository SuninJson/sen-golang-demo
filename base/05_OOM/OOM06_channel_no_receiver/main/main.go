package main

func main() {
	ch := make(chan int)
	// 尝试向无缓冲通道发送数据，但没有 goroutine 接收数据，导致阻塞
	ch <- 1
}
