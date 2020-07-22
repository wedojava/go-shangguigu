package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var myChan chan Person

func main() {
	var lock sync.Mutex
	myChan = make(chan Person, 10)
	for i := 0; i < 10; i++ {
		a := int(rand.Int63n(time.Now().UnixNano()))
		b := Person{"P" + strconv.Itoa(a), a, "A" + strconv.Itoa(a)}
		lock.Lock()
		myChan <- b
		lock.Unlock()
	}
	time.Sleep(10 * time.Second)
	lock.Lock()
	for mc := range myChan {
		fmt.Printf("Name: %s, Age: %d, Address: %s\n", mc.Name, mc.Age, mc.Name)
	}
	lock.Unlock()
}
