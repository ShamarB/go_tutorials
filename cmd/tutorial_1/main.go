package main

import (
	"errors"
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

	numerator := 10
	denominator := 5
	var result, remainder, err = integerDivision(numerator, denominator)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of integer division is %v with a remainder of %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println("Print Me")
	fmt.Println(printValue)
}

func integerDivision(numerator int, denominator int) (int, int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func printMeInt(printValue int) {
	fmt.Println(printValue)
}
