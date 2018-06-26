package main

import (
	"fmt"
)

// 演讲者接口
type Speaker interface {
	// 说
	Say(string)
	// 听
	Listen(string) string
	// 打断、插嘴
	Interrupt(string)
}

// 王兰讲师
type WangLan struct {
	msg string
}

func (this *WangLan) Say(msg string) {
	fmt.Printf("王兰说：%s\n", msg)
}

func (this *WangLan) Listen(msg string) string {
	this.msg = msg
	return msg
}

func (this *WangLan) Interrupt(msg string) {
	this.Say(msg)
}

// 江娄讲师
type JiangLou struct {
	msg string
}

func (this *JiangLou) Say(msg string) {
	fmt.Printf("江娄说：%s\n", msg)
}

func (this *JiangLou) Listen(msg string) string {
	this.msg = msg
	return msg
}

func (this *JiangLou) Interrupt(msg string) {
	this.Say(msg)
}

func main() {
	wl := &WangLan{}
	jl := &JiangLou{}

	var person Speaker
	person = wl
	person.Say("Hello World!")
	person = jl
	person.Say("Good Luck!")
}
