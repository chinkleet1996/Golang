package main

import "fmt"

func modify(s []int) {
	s[1] = 200

}
func main() {
	slice := []int{10, 20, 30, 40}
	fmt.Println("Slice value before modify function call")
	fmt.Println(slice)
	fmt.Println("Slice value after modify function call")
	modify(slice)
	fmt.Println(slice)

}
