package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// b = a // error: cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point) // 类型断言, type assertion
	fmt.Println(b)

	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 //ok
	// y := x.(float32)
	// fmt.Printf("y's type is %T and value is %v", y, y)

	var x interface{}
	var b2 float32 = 1.1
	x = b2 //ok
	// y, ok := x.(float32)
	// y, ok := x.(float64)
	// if ok {
	if y, ok := x.(float32); ok {
		fmt.Println("Convert success.")
		fmt.Printf("y's type is %T and value is %v", y, y)
	} else {
		fmt.Println("Convert fail.")
	}
	fmt.Println("continue...")

}
