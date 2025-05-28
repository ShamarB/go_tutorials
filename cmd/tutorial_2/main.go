package main

import "fmt"

func main() {
	num := 0

	switch num {
	case 0:
		fmt.Println("Number not found")
	case 1:
		fmt.Println("Number found")
	case 2:
		fmt.Println("Number not found")
	}
}
