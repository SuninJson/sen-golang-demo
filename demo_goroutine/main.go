package main

import "fmt"

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	go initIntChan(intChan)

	// 开启4个协程，从intChan中取数，并判断是否为质数，若是将其放入primeChan中
	for i := 0; i < 4; i++ {
		go findPrime(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		close(primeChan)
	}()

	// 打印出质数
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数 = %d\n", res)
	}
}

func initIntChan(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	close(intChan)
}

func findPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeChan <- num
		}
	}

	exitChan <- true
}
