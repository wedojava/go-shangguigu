package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tome jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat
	// get 3rd item in chan have to pull 2 first.
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
}
