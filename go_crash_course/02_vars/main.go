package main

import "fmt"

func main() {
	// using var
	var name string = "Stoney"
	var age int32 = 21
	var isCool = true

	// Shorthand
	//name2 := "Not Marquis"
	//email := "Not stoneyreed25@gmail.com"
	size := 1.3

	name2, email := "Marquis", "stoneyreed25@gmail.com"

	fmt.Println(name, age, isCool, name2, size, email)
	fmt.Printf("%T\n", email)
}
