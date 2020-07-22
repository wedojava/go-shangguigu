package main

import "fmt"

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据=%v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData 读取数据=", v)
	}
	// readData 读取完数据后，任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道，一个放数据，一个放读写状态
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
