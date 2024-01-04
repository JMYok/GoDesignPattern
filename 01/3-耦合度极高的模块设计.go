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

// 此时若要增加driver3，或者driver1也想开宝马

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
