package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	var x1, y int
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil)
	var p1 = f()
	fmt.Println(*p1)
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}
func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++
	return *p
}
