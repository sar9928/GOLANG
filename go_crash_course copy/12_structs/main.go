package main

// Formatter and string converter
import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Alternative person struct
// type Person struct {
// 	firstName, lastName, city, gender string
// 	age                               int
// }

// Greeting method (value reciever) // returns string
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer reciever) // changes struct
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer reciever)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Terry", "Crews", "New York", "m", 21}

	// Alternative
	//person1 := Person{"Samantha", "Smith", "Boston", "f", 25}

	fmt.Println(person1)
	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Reed")
	fmt.Println(person1.greet())
	person2.getMarried("Gozillan")
	fmt.Println(person2.greet())

}
