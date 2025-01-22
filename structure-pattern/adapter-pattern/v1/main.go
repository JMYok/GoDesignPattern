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

// Adaptor 类适配器
type Adaptor struct {
	*Adaptee // 嵌入Adaptee
}

func (a *Adaptor) f1() {
	a.Adaptee.fa()
}

func (a *Adaptor) f2() {
	fmt.Println("f2 re-implemented")
}

// fc不需要重新实现，因为已经通过嵌入Adaptee继承了该方法

func main() {
	adaptor := &Adaptor{&Adaptee{}}
	var target ITarget = adaptor

	target.f1() // 调用fa
	target.f2() // 调用重写后的f2
	target.fc() // 调用fa中的fc
}
