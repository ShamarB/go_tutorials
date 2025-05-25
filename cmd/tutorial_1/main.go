package main

import (
	"fmt"
)

func main() {
	var intNum int = 28
	fmt.Println(intNum)

	myVar := "text"
	fmt.Println(myVar)

	var mybool bool = true
	fmt.Println(mybool)

	const pi float64 = 3.14159
	fmt.Println(pi)

	printMe("Hello")

	numerator := 11
	denominator := 2
	var result, remainder int = integerDivision(numerator, denominator)
	fmt.Printf("The result of integer division is %v with a remainder of %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println("Print Me")
	fmt.Println(printValue)
}

func integerDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}

func printMeInt(printValue int) {
	fmt.Println(printValue)
}
