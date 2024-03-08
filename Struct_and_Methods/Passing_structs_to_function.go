package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func CalcArea(c *Circle) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	(*c).area = area
}

func main() {
	c := Circle{x: 7, y: 9, radius: 5, area: 0}
	fmt.Println("struct field value before function call...")
	fmt.Printf("%+v \n", c)
	// calling the CalcArea function by passing struct variable c.
	CalcArea(&c)
	fmt.Println("struct field values after function call...")
	fmt.Printf("%+v \n", c)
}
