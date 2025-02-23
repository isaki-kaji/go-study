package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "もぐらの穴", "montain", "タコの足は8本でイカの足は10本"}

	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Println("短い単語だ")
		case 6:
			fmt.Println("ちょうど良い長さだ")
		case 7, 8:
		default:
			fmt.Println("よく分からない")
		}
	}
}
