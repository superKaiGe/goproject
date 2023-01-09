package main

import (
	"fmt"
)

func main() {
	var a animal
	var c cat
	a = c
	a.printInfo()
	var d dog
	a = d
	a.printInfo()
}

type animal interface {
	printInfo()
}
type cat int
type dog int

func (c cat) printInfo() {
	fmt.Println("a cat")
}
func (d dog) printInfo() {
	fmt.Println("a dog")
}
