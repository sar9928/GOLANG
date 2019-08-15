package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr = [2]string{"banana", "kiwi"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Pineapple", "apple", "grape"}

	// = 3
	fmt.Println(len(fruitSlice))
	// = pineapple, apple  // starts at 0 and ends at 2
	fmt.Println(fruitSlice[0:2])
}
