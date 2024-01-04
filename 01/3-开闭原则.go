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

// 新加功能直接扩展
type TransferBanker struct {
	AbstractBanker
}

func (s *TransferBanker) DoBusiness() {
	fmt.Println("转账业务")
}
