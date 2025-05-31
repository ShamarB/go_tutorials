package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'my string' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nMy Rune: %v", myRune)

	//when concatenating you're builing a new string which is inefficient, instead import string builder
	var stringSlice = []string{"s", "o", "m", "e"}
	var strBuilder strings.Builder
	for i := range stringSlice {
		strBuilder.WriteString(stringSlice[i])
	}
	var catString = strBuilder.String()
	fmt.Printf("\n%v", catString)
	//catString[0] = 'a' --> String are immutable in go

}
