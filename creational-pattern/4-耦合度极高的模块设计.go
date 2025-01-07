package main

import "fmt"

// 司机driver1,drive2 汽车:奔驰,宝马
// 需求1开1,2开2

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW is running")
}

type Driver1 struct {
}

func (d *Driver1) Drive(benz *Benz) {
	fmt.Println("driver1 drive benz")
	benz.Run()
}

type Driver2 struct {
}

func (d *Driver2) Drive(bmw *BMW) {
	fmt.Println("driver2 drive bmw")
	bmw.Run()
}

// 此时若要增加driver3，或者driver1也想开宝马，此时就要增加模块
// 而且增加的模块（例如新增车 卡宴）需要和之前每个模块进行调试（调试driver1，driver2的drive方法），
//确保如果driver1和driver2想开卡宴时，功能是正确的
// 所以导致各个模块的耦合度过高，如果系统复杂，都不敢修改或者新增功能，
// 因为不知道新增的模块需要和哪些模块交互，修改的模块会影响哪些模块，除非非常懂业务。
// 依赖转置原则就是为了解决这个问题，将系统设计为业务逻辑层、抽象层和实现层。

type Driver3 struct {
}

func (d *Driver3) Drive(bmw *BMW) {
	fmt.Println("driver3 drive bmw")
	bmw.Run()
}

//

func main() {
	driver1 := &Driver1{}
	driver2 := &Driver2{}
	benz := &Benz{}
	bmw := &BMW{}

	driver1.Drive(benz)
	driver2.Drive(bmw)
}
