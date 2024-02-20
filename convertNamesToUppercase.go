package main

import (
	"fmt"
	"strings"
)

func mapStrings(fn func(string) string, list []string) []string {
	mappedList := make([]string, len(list))
	for i, value := range list {
		mappedList[i] = fn(value)

	}
	return mappedList
}

func main() {
	names := []string{"joe", "harry", "chloe", "suzaine"}
	toUpper := func(s string) string { return strings.ToUpper(s) }
	upper_Names := mapStrings(toUpper, names)
	fmt.Println("Here are the names with capital letters: \n", upper_Names)
}
