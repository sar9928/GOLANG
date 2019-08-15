package main

import "fmt"

func main() {
	// Defined map
	emails := make(map[string]string)

	// Assign kv "key value"
	emails["Bob"] = "bob@gmail.com"
	emails["Stoney"] = "stoney@gmail.com"
	emails["Gary"] = "gary@gmail.com"
	emails["Stacy"] = "stacy@gmail.com"

	fmt.Println("Old emails:", emails)
	fmt.Println("My email:", emails["Stoney"])
	fmt.Println("Email length:", len(emails))

	delete(emails, "Bob")
	fmt.Println("New emails:", emails)

	// Declare map and add kv "Key Value"
	superEmails := map[string]string{"Bruce": "batman@gmail.com", "Clark": "superman@gmail.com"}

	superEmails["Stoney"] = "superstone@gmail.com"
	fmt.Println(superEmails)

}
