package main

import "fmt"

func reverseArray(arr []int) []int {
	n := len(arr)

	reversed := make([]int, n)

	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}

	return reversed
}

func main() {
	original := []int{10, 20, 30, 40, 50}

	fmt.Println("Original array: ", original)

	reversed := reverseArray(original)

	fmt.Println("Reversed array:", reversed)
}
