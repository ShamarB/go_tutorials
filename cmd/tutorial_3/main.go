package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(intArr[1:3])

	//Get memory location of element
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length of slice is %v with capacity of %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length of slice is %v with capacity of %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	//Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 28, "Sarah": 39}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Sarah"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Ivalid name")
	}

	//delete(myMap2, "Adam")

	//Loops
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	//indexing over arrays or slices
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	//while loop equivalent
	var i int = 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
