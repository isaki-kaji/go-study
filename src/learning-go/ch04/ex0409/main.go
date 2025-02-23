package main

import "fmt"

func main() {
	i := 1

	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	fmt.Println("finish.")

	j := 2
	for {
		fmt.Println(j)

		if j%2 == 0 {
			break
		}
	}
}
