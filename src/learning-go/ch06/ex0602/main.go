package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

// 　基本型のリテラルや定数に「&」はつけらrない
func stringp(s string) *string {
	return &s
}

func main() {

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p)
	fmt.Println(*p.MiddleName)
}
