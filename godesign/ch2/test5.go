package main

import "fmt"

func main() {
	s := "abc"
	b := []byte(s)
	fmt.Println("%v", b)
	s2 := string(b)
	fmt.Println(s2)
}
