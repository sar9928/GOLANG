package main

import "fmt"

func main() {
	ids := []int{33, 84, 23, 90, 12, 69, 6}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum of all ids =", sum)

	// Range with map
	superEmails := map[string]string{"Bruce": "batman@gmail.com", "Clark": "superman@gmail.com"}

	// Key and value pair
	for key, value := range superEmails {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Only key
	for key := range superEmails {
		//fmt.Printf("Name: %s\n", key)
		fmt.Println("Name:", key)
	}
}
