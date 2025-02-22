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

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) //5 true

	v, ok = m["world"]
	fmt.Println(v, ok) // 0 true

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 true
}
