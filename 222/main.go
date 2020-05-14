package main

import "fmt"

type Usb interface {
	/* TODO: add methods */
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func main() {
	var usbs [3]Usb // polymorphic array
	usbs[0] = Phone{"vivo"}
	usbs[1] = Phone{"iPhone"}
	usbs[2] = Camera{"Nikon"}
	fmt.Println(usbs)
}
