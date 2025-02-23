package main

import "fmt"

func main() {
	uniqueNames := map[string]bool{
		"花子": true,
		"太郎": true,
		"次郎": true,
	}

	for k := range uniqueNames {
		fmt.Println(k)
	}
}
