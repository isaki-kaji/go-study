package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("小さいです。")
	} else if n > 5 {
		fmt.Println("大きいです。")
	} else {
		fmt.Println("いい感じです。")
	}
}
