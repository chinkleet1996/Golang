package main

import "fmt"

func main() {
	s := "Penetrationtesting"
	fmt.Printf("%T %v \n", s, s)
	ps := &s
	fmt.Println("------------------------------")
	fmt.Println("The memory address of the pointer which holds the address of the variable s:", ps)
	fmt.Println("-------------------------------")
	*ps = "Metasploithacking"
	fmt.Println("The value after dereferencing with the pointer")
	fmt.Println("The memory address of the pointer ps after dereferencing the value of the variable s:", ps)
	fmt.Printf("%T %v \n", s, s)
}
