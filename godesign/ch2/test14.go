package main

import "fmt"

func main() {
	r := [...]int{99: -1}
	fmt.Println(r)
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	//d := [3]int{1, 2}
	//fmt.Println(a == d)
}
