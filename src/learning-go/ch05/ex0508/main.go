package main

import (
	"errors"
	"fmt"
)

// 使うべきではない
func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("0で割ることはできません")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func main() {
	fmt.Println(divAndRemainder(4, 4))
}
