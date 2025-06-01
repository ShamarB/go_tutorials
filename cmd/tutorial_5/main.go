package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	//owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// type owner struct {
// 	name string
// }

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You need to fill up first")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{20, 15} //owner{"Shamar"}}
	myEngine.mpg = 25
	fmt.Println(myEngine.mpg, myEngine.gallons) //, myEngine.name)

	//fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	var myTesla electricEngine = electricEngine{25, 15}
	canMakeIt(myTesla, 77)
}
