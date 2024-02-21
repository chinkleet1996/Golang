package main

import (
	"fmt"
	"math"
)

type CircleFunc func(float64) float64

func main() {
	radius := getInput("Enter the radius of the circle: ")

	area := calculateArea(radius)
	perimeter := calculatePerimeter(radius)
	diameter := calculateDiameter(radius)

	fmt.Println("---------------------------")
	fmt.Println("Area of the circle is:", area)
	fmt.Println("Perimeter of the circle is:", perimeter)
	fmt.Println("Diameter of the circle is:", diameter)
}

func getInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input

}

func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func calculatePerimeter(radius float64) float64 {
	return 2 * math.Pi * radius

}

func calculateDiameter(radius float64) float64 {
	return 2 * radius

}
