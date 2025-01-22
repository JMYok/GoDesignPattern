package main

import (
	"fmt"
)

// 定义目标接口
type ITarget interface {
	f1()
	f2()
	fc()
}

// Adaptee 类
type Adaptee struct{}

func (a *Adaptee) fa() {
	fmt.Println("fa called")
}

func (a *Adaptee) fb() {
	fmt.Println("fb called")
}

func (a *Adaptee) fc() {
	fmt.Println("fc called")
}

// Adaptor 对象适配器
type Adaptor struct {
	adaptee *Adaptee
}

func NewAdaptor(adaptee *Adaptee) *Adaptor {
	return &Adaptor{adaptee: adaptee}
}

func (a *Adaptor) f1() {
	a.adaptee.fa()
}

func (a *Adaptor) f2() {
	fmt.Println("f2 re-implemented")
}

func (a *Adaptor) fc() {
	a.adaptee.fc()
}

func main() {
	adaptee := &Adaptee{}
	adaptor := NewAdaptor(adaptee)
	var target ITarget = adaptor

	target.f1() // 调用fa
	target.f2() // 调用重写后的f2
	target.fc() // 调用fa中的fc
}
