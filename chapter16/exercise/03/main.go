// 1. 启动一个协程，将1-2000的数放入到一个channel中，比如numChan
// 2. 启动8个协程，从numChan取出数(比如n)，并计算 1+...+n的值，并存放到resChan
// 3. 最后8个协程协同完成工作后，再遍历resChan，显示结果
// 4. 考虑 resChan chan int 是否合适
package main

var (
	numChan chan int
	resChan chan map[int]int
)

func save() {
	for i := 0; i < 2001; i++ {
		numChan <- i + 1
	}
	close(numChan)
}

func accumulate(n int, resChan chan map[int]int) {
	var res map[int]int
	a := 0
	nums := <-numChan
	for i := 0; i < nums; i++ {
		a += i + 1
	}
	res[n] = a
	resChan <- res
	close(resChan)
}

func treat() {
	for i := 0; i < 8; i++ {
		go accumulate(i+1, resChan)
	}
}

func main() {

}
