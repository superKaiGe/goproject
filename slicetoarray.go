package main

import (
	"fmt"
)

func main() {
	a := [3]int{11, 12, 13}
	b := a[:]
	b[1] = b[1] + 10
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	//b = []int{11, 12, 13}
	//var p = (*[3]int)(unsafe.Pointer(&b[0]))
	//p[1] += 10
	//fmt.Printf("%v\n,%t", b, b)
	fmt.Println("---------------")
	b = []int{11, 12, 13}
	p := (*[3]int)(b)
	p[1] = p[1] + 10
	fmt.Printf("%v\n", b)
}
