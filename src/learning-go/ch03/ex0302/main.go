package main

import "fmt"

func main() {
	x := make([]int, 0, 10) //スライスを長さ0で初期化してからappendしていくのが良さそう
	fmt.Println(x)

	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)
}
