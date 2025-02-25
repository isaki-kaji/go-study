package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			continue
		}

		op := expression[1]

		//　キーから存在チェックができる
		opFunc, ok := opMap[op]
		if !ok {
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
