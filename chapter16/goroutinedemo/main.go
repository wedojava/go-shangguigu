package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main: hello world", strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}

}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test: hello world", strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
