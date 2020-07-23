package main

import "fmt"

func main() {
	// 使用 select 可以解决从管道取数据阻塞的问题

	// 1. decelare channel with 10 int data
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2. decelare channel with 5 string data
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 实际开发中，可能我们不好确定什么时候关闭管道。
	// 使用 select 来遍历
	// label: // 不推荐用标签来break退出循环
	for {
		select {
		// 注意！如果管道一直没有关闭也不会阻塞而死锁
		// 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Println("get data from intChan: ", v)
		case v := <-strChan:
			fmt.Println("get data from strChan: ", v)
		default:
			fmt.Println("... got nothing!")
			return
			// break label
		}
	}
}
