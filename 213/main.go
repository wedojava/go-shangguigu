package main

import "fmt"

type Computer struct {
}

type Phone struct {
}

type Camera struct {
}

type IUsb interface {
	Start()
	Stop()
}

func (p Phone) Start() {
	fmt.Println("Phone start")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop")
}

func (c Camera) Start() {
	fmt.Println("Camera start")
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
