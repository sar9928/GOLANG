package main

import "fmt"

func main() {
	// int
	a := 5
	// pointer 
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
