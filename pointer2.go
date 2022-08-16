package main

import "unsafe"

type foo struct {
	id   string
	age  int8
	addr string
}

func main() {
	var p1 *int
	var p2 *bool
	var p3 *byte
	var p4 *[20]int
	var p5 *foo
	var p6 unsafe.Pointer
	println(unsafe.Sizeof(p1))
	println(unsafe.Sizeof(p2))
	println(unsafe.Sizeof(p3))
	println(unsafe.Sizeof(p4))
	println(unsafe.Sizeof(p5))
	println(unsafe.Sizeof(p6))
}
