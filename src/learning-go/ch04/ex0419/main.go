package main

import "fmt"

func main() {
	samples := []string{"hello", "これは感じ文字列"}

outer:
	for _, sample := range samples {
		for i, v := range sample {
			fmt.Println(i, v, string(v))
			if v == 'は' {
				continue outer
			}
		}
		fmt.Println()
	}
}
