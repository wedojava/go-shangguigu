package main

import "fmt"

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "can climbing trees.")
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "studied flying")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "studied swimming!")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
