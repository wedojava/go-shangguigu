package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	// 看看 intChan 是什么
	fmt.Printf("intChan: %v\n", intChan)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// 不能超过管道容量
	// intChan <- 50

	// 查看管道长度和容量
	fmt.Printf("intChan: cap=%v, len=%v\n", cap(intChan), len(intChan))

	// 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("intChan: cap=%v, len=%v\n", cap(intChan), len(intChan))

	// 如果管道数据全部取完再取会出现 deadlock 错误。
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan // err: fatal error: all goroutines are asleep - deadlock!

	fmt.Println("num3=", num3, "num4=", num4)
	fmt.Println(num5)
}
