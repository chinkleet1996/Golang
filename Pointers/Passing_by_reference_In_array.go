package main

import "fmt"

// Function modify takes a pointer to an array of integers as a parameter and subtracts 5 from each element.
func modify(numbers *[3]int) {
	// Iterate over the array pointed to by numbers using a range loop.
	for i := range numbers {
		// Subtract 5 from each element of the array.
		numbers[i] -= 5
	}
}

func main() {
	// Initialize an array named arr with integer values 10, 20, and 30.
	arr := [3]int{10, 20, 30}
	// Print the original array arr.
	fmt.Println(arr)

	// Call the modify function with a pointer to the array arr.
	modify(&arr)

	// Print the array arr after modifying its elements.
	fmt.Println(arr)
}
