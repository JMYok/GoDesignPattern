package main

import "fmt"

// 扩展类可以，但不修改原来的类

type AbstractBanker interface {
	DoBusiness()
}

type SaveBanker struct {
	AbstractBanker
}

func (s *SaveBanker) DoBusiness() {
	fmt.Println("存款业务")
}

// TransferBanker （新加功能直接扩展）
type TransferBanker struct {
	AbstractBanker
}

type ShareBanker struct {
	AbstractBanker
}

func (t *TransferBanker) DoBusiness() {
	fmt.Println("转账业务")
}

func (s *ShareBanker) DoBusiness() {
	fmt.Println("股票业务")
}

// bankBusiness 该模式也可以抽象出一个架构层
func bankBusiness(banker AbstractBanker) {
	//利用多态，父指针能够调到子对象的方法
	banker.DoBusiness()
}

func main() {
	//股票业务
	bankBusiness(&ShareBanker{})
	//转账业务
	bankBusiness(&TransferBanker{})
}
