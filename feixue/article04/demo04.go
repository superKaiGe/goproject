package main

import "fmt"

type animal interface {
	printInfo()
}
type cat int

func (c *cat) printInfo() {
	fmt.Println("a cat")
}
func invoke(a animal) {
	a.printInfo()
}
func main() {
	var c cat
	//invoke(c)
	//指针作为参数传递
	invoke(&c)
}
