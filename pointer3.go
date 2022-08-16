package main

import "fmt"

func main() {
	var a int = 17
	var p *int = &a
	//var p1* = &p
	println(*p)
	(*p) += 3
	(*p) += 15

	println(a)
	println(*p)
	println(p)
	println(&p)
	fmt.Printf("%p\n", p)
}
