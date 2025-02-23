package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	if n := rand.Intn(10); n == 0 {
		fmt.Println("小さすぎます。")
	} else {
		fmt.Println("悪くないです。")
	}
}
