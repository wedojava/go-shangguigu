package main

import "fmt"

type Computer struct {
	rate float64
}

type Phone struct {
	Num int
}

type Camera struct {
	Name string
}

type IUsb interface {
	Start()
	Stop()
}

func (p Phone) Start() {
	fmt.Println("Phone start", p.Num)
}

func (p Phone) Stop() {
	fmt.Println("Phone stop")
}

func (c Camera) Start() {
	fmt.Println("Camera start", c.Name)
}

func (c Camera) Stop() {
	fmt.Println("Camera stop")
}

func (c Computer) Working(usb IUsb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	p := Phone{}
	c := Camera{}
	computer.Working(p)
	computer.Working(c)
}
