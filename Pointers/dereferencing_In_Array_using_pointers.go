package main

import "fmt"

func main() {
	// Define an array y with 3 integer elements
	y := [3]int{10, 20, 30}

	// Print the initial value of array y
	fmt.Printf("%v \n", y)

	// Change the first element of array y indirectly using pointer operations
	(*&y)[0] = 100

	// Print a separator
	fmt.Println("--------")

	// Print the modified value of array y
	fmt.Printf("%v \n", y)
}
