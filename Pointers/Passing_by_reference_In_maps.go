package main

import "fmt"

func modify(m map[string]int) {
	m["C"] = 75

}
func main() {
	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["B"] = 70
	fmt.Println("Map value before modify function call")
	fmt.Println(ascii_codes)
	fmt.Println("Map value after modify function call")
	modify(ascii_codes)
	fmt.Println(ascii_codes)

}
