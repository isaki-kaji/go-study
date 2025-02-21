package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x == nil) //true

	x = append(x, 10)
	fmt.Println(x)

	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x)
}
