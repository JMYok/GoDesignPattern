package main

import "fmt"

type Banker struct{}

func (b *Banker) Save() {
	fmt.Println("存款业务")
}

func (b *Banker) Transfer() {
	fmt.Println("转账业务")
}

func (b *Banker) Pay() {
	fmt.Println("支付业务")
}

// 新加功能,对原有类进行修改
// 不知道新加功能会不会和其他功能有耦合，影响原来类功能的使用
func (b *Banker) Shares() {
	fmt.Println("股票业务")
}

func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()
}
