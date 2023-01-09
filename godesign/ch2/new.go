package main

import "fmt"

func main() {
	//p := 1
	//fmt.Println(&p)
	//x := &p
	//fmt.Println(*x)
	//*x = 2
	//fmt.Println(p)
	p := new(int)
	q := new(int)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(p == q)
	var x *int
	fmt.Println(*(&x))
}
