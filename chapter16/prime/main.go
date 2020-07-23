package main

import "fmt"

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i + 1
	}
	close(intChan)
}

func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		flag = true
		// if it is prime number
		for i := 2; i < num; i++ {
			if num%i == 0 { // num is not prime number
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
	fmt.Println("a primeNum goroutine done.")
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 8000)
	exitChan := make(chan bool, 4)

	// 开启一个 goroutine 向 intChan 放入 1-8000
	go putNum(intChan)
	// 开启4个 goroutine 从 intChan 取出数据，并判断是否为 prime 然后放入 primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 遍历获得结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("prime=%d\n", res)
	}
	fmt.Println("done.")
}
