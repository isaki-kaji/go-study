package main

import "fmt"

func main() {
	fmt.Println(div(3, 2))
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}
