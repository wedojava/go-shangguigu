package main

import (
	"fmt"

	"github.com/wedojava/shangguigu/chapter10/factory/model"
)

func main() {
	// Create a Student implement
	// var stu = model.student{
	//         Name:  "Tom",
	//         Score: 78.9,
	// }

	var stu = model.NewStudent("tom", 78.9)
	fmt.Println(*stu)
	fmt.Println("name is", stu.Name)
}
