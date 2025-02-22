package main

import "fmt"

func main() {
	totalwins := map[string]int{}
	totalwins["ライターズ"] = 1
	totalwins["ナイツ"] = 2
	fmt.Println(totalwins["ライターズ"])
	fmt.Println(totalwins["ミュージシャンズ"])
	totalwins["ミュージシャンズ"]++
	fmt.Println(totalwins["ミュージシャンズ"]) // 1
	totalwins["ナイツ"] = 3
	fmt.Println(totalwins["ナイツ"])
}
