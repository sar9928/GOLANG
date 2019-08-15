package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// scanner for input
	scanner := bufio.NewScanner(os.Stdin)

	tries := 1
	low := 1
	high := 100
	lowCheck := high
	highCheck := low

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	// checks for ENTER input
	scanner.Scan()

	for {
		// Utilize binary search
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("(a) Is that too high?")
		fmt.Println("(b) Is that too low?")
		fmt.Println("(c) Is that correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			if high < highCheck {
				fmt.Println("Hey stop cheating I asked for a number between", highCheck, "and", lowCheck)
				break
			}
			tries++
		} else if response == "b" {
			low = guess + 1
			if low > lowCheck {
				fmt.Println("Hey stop cheating I asked for a number between", highCheck, "and", lowCheck)
				break
			}
			tries++
		} else if response == "c" {
			fmt.Println("I won in only", tries, "tries")
			break
		} else {
			fmt.Println("Invalid response, try again")
		}

	}

}
