package main

import "fmt"

// Creating interface Student with two methods
type Student interface {
	getPercentage() int
	getName() string
}

// creating struct with name and grades as fields
type Undergrad struct {
	name   string
	grades []int
}

// function to implement getPercentage method of the interface
func (u Undergrad) getPercentage() int {
	sum := 0
	for _, value := range u.grades {
		sum += value
	}

	return ((sum * 100) / (len(u.grades) * 100))

}

// function to implement getName method of the interface
func (u Undergrad) getName() string {
	return u.name
}

// function to call both the methods with Student interface object s
func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	// initialization of the Undergrad struct with name and grades
	u := Undergrad{"Ross", []int{98, 76, 85}}
	printData(u)
}
