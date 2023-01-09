package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for k, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(k, arg)
	}
	a := os.Args[0]
	fmt.Println(a)
	fmt.Println(s)
}
