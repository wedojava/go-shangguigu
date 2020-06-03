package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

// 1. B 的对象可以使用 A 里面所有的对象和方法, 不管它是否是私有
type B struct {
	A    // 嵌入匿名结构体
	Name string
}

func (a *A) SayOK() {
	fmt.Println("A said OK!", a.Name)
}

func (a *A) hello() {
	fmt.Println("A said hello", a.Name)
}

func (b *B) SayOK() {
	fmt.Println("B said OK!", b.Name)
}

func main() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOK() // A said OK! tom
	b.A.hello() // A said hello tom

	// 2. 匿名结构体的访问可以简化
	b.Name = "tom2"
	b.age = 19
	b.SayOK() // B said OK! tom2
	b.hello() // A said hello tom
	// 3. ↑ 编译器会先看 b 是否有相应字段和方法,如果没有就去看 b 中嵌入的匿名结构体是否有相应的字段和方法,然后调用
	// 4. 调用哪个字段和方法遵循就近原则
	b.Name = "jack"
	b.age = 100
	b.SayOK() // B said OK! jack
	b.hello() // A said hello tom
}
