package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func main() {
	mov := getMovie("xyz", 2.1)
	fmt.Printf("Printing the mov object value %+v \n", mov)
	mov1 := getMovie("abc", 3.3)
	fmt.Printf("Printing the mov1 object value %+v \n", mov1)
	// Initializing an empty slice with name movies..
	movies := make([]Movie, 5)
	// appending both mov and mov1 object value to movies slice
	movies = append(movies, mov)
	movies = append(movies, mov1)
	// Looping through the slice and storing the
	for _, value := range movies {
		fmt.Println(value)
	}
}
