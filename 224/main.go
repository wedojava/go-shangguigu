package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (p Phone) Start() {
	fmt.Println("Phone starting.")
}
func (p Phone) Stop() {
	fmt.Println("Phone stoping.")
}

func (c Camera) Start() {
	fmt.Println("Camera starting.")
}
func (c Camera) Stop() {
	fmt.Println("Camera stoping.")
}

func (p Phone) Call() {
	fmt.Println("Phone calling ...")
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	// type assertion
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbs [3]Usb // polymorphic array
	usbs[0] = Phone{"vivo"}
	usbs[1] = Phone{"iPhone"}
	usbs[2] = Camera{"Nikon"}
	// fmt.Println(usbs)
	var computer Computer
	for _, usb := range usbs {
		computer.Working(usb)
	}
}
