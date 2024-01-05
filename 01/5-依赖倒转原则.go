package main

import "fmt"

//--------->抽象层<---------

type Driver interface {
	Drive(car Car)
}

type Car interface {
	Run()
}

//--------->实现层<---------
// 依赖抽象层，面向接口开发

type Driver5 struct {
	//..
}

type Driver6 struct {
	//..
}

func (d *Driver5) Drive(car Car) {
	fmt.Println("driver5 is driving...")
	car.Run()
}

func (d *Driver6) Drive(car Car) {
	fmt.Println("driver6 is driving...")
	car.Run()
}

type BenZ struct {
	//..
}

type Keyan struct {
	//..
}

func (b *BenZ) Run() {
	fmt.Println("BenZ is running...")
}

func (k *Keyan) Run() {
	fmt.Println("Keyan is running")
}

// --------->业务逻辑层<---------
// 依赖抽象层,面向接口开发
// 这样我之后在新增一个模块，我只需要依赖接口，不需要调试其他模块的功能，只需要调试当前模块功能。
// 例如新增了丰田车，实现Car接口，只需要确定Run是否正常就行，其他Driver模块依赖了Drive，不需要调试了。
func main() {
	var d5 Driver
	d5 = new(Driver5)
	var bz Car
	bz = new(BenZ)
	d5.Drive(bz)

	var d6 Driver
	d6 = new(Driver6)
	var ky Car
	ky = new(Keyan)
	d6.Drive(ky)
}
