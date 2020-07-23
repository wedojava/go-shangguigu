// 1. 启动一个协程，将1-2000的数放入到一个channel中，比如numChan
// 2. 启动8个协程，从numChan取出数(比如n)，并计算 1+...+n的值，并存放到resChan
// 3. 最后8个协程协同完成工作后，再遍历resChan，显示结果
// 4. 考虑 resChan chan int 是否合适
package main

import (
	"fmt"
)

var (
	numChan  chan int
	resChan  chan map[int]int
	exitChan chan bool
)

func save(numCh chan int) {
	for i := 0; i < 2000; i++ {
		numCh <- i + 1
	}
	close(numCh)
}

func accumulate(srcCh chan int, desCh chan map[int]int) {
	for {
		res := make(map[int]int)
		a := 0
		v, ok := <-srcCh
		if !ok {
			break
		}
		fmt.Println("accumulate get nums: ", v)
		for i := 0; i < v; i++ {
			a += i + 1
		}
		res[v] = a
		desCh <- res
	}
	exitChan <- true
}

func treat(n int, srcCh chan int, desCh chan map[int]int) {
	for i := 0; i < n; i++ {
		go accumulate(srcCh, desCh)
	}
}

func printCh(desCh chan map[int]int) {
	for des := range desCh {
		fmt.Println(des)
	}
}

func main() {
	numChan = make(chan int, 2000)
	resChan = make(chan map[int]int, 2000)
	exitChan = make(chan bool, 1)
	// 1. 启动一个协程，将1-2000的数放入到一个channel中，比如numChan
	go save(numChan)
	// 2. 启动8个协程，从numChan取出数(比如n)，并计算 1+...+n的值，并存放到resChan
	treat(8, numChan, resChan)
	// time.Sleep(10 * time.Second)
	i := 0
	for {
		done, ok := <-exitChan
		if !ok {
			continue
		}
		if done {
			i++
		}
		if i == 8 {
			break
		}
	}
	close(exitChan)
	close(resChan)
	// 3. 最后8个协程协同完成工作后，再遍历resChan，显示结果
	printCh(resChan)
}
