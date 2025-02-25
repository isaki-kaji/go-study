package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ファイルが指定されていません")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	doSomething("a", "b")
}

// 名前付き戻り値を使うことによってdefer内でerrが使える
func doSomething(value1, value2 string) (err error) {
	err = errors.New("エラーです")

	defer func() {
		if err == nil {
			fmt.Println()
		}
		if err != nil {
			fmt.Println(value1, value2)
		}
	}()

	return nil
}
