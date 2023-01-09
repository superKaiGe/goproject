package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	var runes1 []rune
	runes1 = []rune("Hello,世界")
	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n", runes1)
	fmt.Printf("%T\n", runes1)
}
