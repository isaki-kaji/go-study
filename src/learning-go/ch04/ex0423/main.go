package main

func main() {
	words := []string{"hi", "wasshoi matsu", "hello"}
	for _, v := range words {
		switch wordlen := len(v); {
		case wordlen < 5:
			println("short")
		case wordlen > 10:
			println("long")
		default:
			println("medium")
		}
	}
}
