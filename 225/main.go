package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for i, v := range items {
		i++
		switch v := v.(type) {
		case bool:
			fmt.Printf("para %v is bool, value is %v\n", i, v)
		case float32:
			fmt.Printf("para %v is float32, value is %v\n", i, v)
		case float64:
			fmt.Printf("para %v is float64, value is %v\n", i, v)
		case int, int32, int64:
			fmt.Printf("para %v is int(32/64), value is %v\n", i, v)
		case string:
			fmt.Printf("para %v is string, value is %v\n", i, v)
		case Student:
			fmt.Printf("para %v is Student, value is %v\n", i, v)
		case *Student:
			fmt.Printf("para %v is *Student, value is %v\n", i, v)
		default:
			fmt.Printf("para %v is unknow type, value is %v\n", i, v)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int32 = 30
	var name string = "Tomy"
	address := "DC"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, n4, address, name, stu1, stu2)

}
