package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// intChan <- 300 // err occur
	fmt.Println("done.")
	n1 := <-intChan
	fmt.Println(n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 遍历前关闭Channel则会正常遍历数据，最后退出遍历,如果不关闭则会出现 deadlock 错误。
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
