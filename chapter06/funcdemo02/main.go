package main

import (
	"fmt"
)

func main() {
	n1 := 10
	test(n1)
	fmt.Println(n1)
	sum := getSum(n1, n1)
	fmt.Println("main sum: ", sum)
	for i := 0; i < 10; i++ {
		fmt.Printf("--")
	}
	fmt.Println()
	recursive(6)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("n1 = ", n1)
}

func getSum(n1, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum: ", sum)
	return sum
}

func recursive(n int) {
	if n > 2 {
		n--
		recursive(n)
	}
	fmt.Println("n = ", n)
}
