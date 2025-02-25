package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Pattn", 32},
		{"Zatt", "J", 18},
		{"T", "Wada", 42},
	}

	sort.Slice(people, func(i int, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	fmt.Println(people)
}
