package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) incLength(n int) {

	for i := 0; i < n; i++ {
		r.length += i
	}
}

func main() {
	r := Rectangle{breadth: 10, length: 5}
	// Printing the area of the rectangle before lenght increased...
	fmt.Println(r.area())
	// Printing the struct Rectangle value
	fmt.Println(r)
	// calling the incLenght method with 7 as value passed
	r.incLength(7)
	// Printing the area  of the rectangle after length increased
	fmt.Println(r.area())
	// Printing struct Rectangle value after length increased
	fmt.Println(r)
}
