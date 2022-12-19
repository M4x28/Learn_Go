package main

import (
	"fmt"
)

func main() {
	// Create an empty slice of floats
	numbers := []float64{}

	// Print the numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Print("Enter a number: ")

		fmt.Scan(&name)

		// Print a greeting
		fmt.Printf("Hello, %s!\n", name)
	}
}
