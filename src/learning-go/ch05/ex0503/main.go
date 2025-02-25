package main

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(2, 4, 5))
	fmt.Println(addTo(2, []int{3, 4, 5}...))
}
