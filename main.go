package main

import "fmt"

func reduceIntegers(fn func(int, int) int, list []int, initialValue int) int {
	result := initialValue
	for _, value := range list {
		result = fn(result, value)
	}

	return result
}

func main() {
	numbers := []int{1286, 2345, 5600, 6756, 3344}
	sum := func(x, y int) int { return x + y }
	total := reduceIntegers(sum, numbers, 0)
	fmt.Println("The total sales revenue for this month is: ", total)

}
