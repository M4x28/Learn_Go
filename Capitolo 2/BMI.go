package main

import (
	"fmt"
)

func main() {
	// Prompt the user for their weight and height
	var weight, height float64
	fmt.Print("Enter your weight (in kg): ")
	fmt.Scan(&weight)
	fmt.Print("Enter your height (in meters): ")
	fmt.Scan(&height)

	// Calculate the BMI
	bmi := weight / (height * height)

	// Print the BMI
	fmt.Printf("Your BMI is %.2f\n", bmi)
}
