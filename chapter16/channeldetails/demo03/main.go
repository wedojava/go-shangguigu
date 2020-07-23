package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello world!")
	}
}

func test() {
	// 这里我们可以使用 defer + reconver 来解决
	var myMap map[int]string
	myMap[0] = "golang" // error, because no make for map
}

func main() {
	go sayHello()
	go test()
	time.Sleep(11 * time.Second)
}
