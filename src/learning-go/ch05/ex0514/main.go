package main

import (
	"fmt"
	"log"
	"os"
)

// closerを呼ぶのを強制できる
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return file, func() { file.Close() }, err
}

func main() {
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	fmt.Println(f)
}
