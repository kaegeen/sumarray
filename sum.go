package main

import (
	"fmt"
)

// sumArray calculates the sum of an array of integers.
func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	// Example array of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Calculate the sum
	result := sumArray(numbers)

	// Print the result
	fmt.Printf("The sum of the array is: %d\n", result)
}
