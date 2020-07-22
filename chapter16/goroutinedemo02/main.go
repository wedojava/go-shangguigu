package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mymap = make(map[int]int, 10)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 0; i < n; i++ {
		res *= i + 1
	}
	lock.Lock()
	mymap[n] = res
	lock.Unlock()
}

func main() {
	for i := 0; i < 200; i++ {
		go test(i + 1)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(mymap)
}
