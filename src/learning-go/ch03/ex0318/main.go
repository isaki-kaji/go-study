package main

import "fmt"

type firstPerson struct {
	name     string
	age      int
	children []string
}

type secondPerson struct {
	name     string
	age      int
	children []string
}

func main() {
	kaji := firstPerson{
		"isaki",
		28,
		[]string{"tara", "sira"},
	}

	kaji2 := secondPerson(kaji)

	fmt.Println(kaji2)
}
