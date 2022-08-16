package main

import (
	"fmt"
	"unsafe"
)

type T int

func main() {
	var p *int
	var p1 = unsafe.Pointer(p) //任意指针显示转换为unsafe.Pointer
	fmt.Println(p1)
	p = (*int)(p1) //unsafe.Pointer也可以显示转换为任意指针类型
	fmt.Println(p)
	//var b byte = 10
	//println(p == nil)
	//var a int = 13
	//p = &a
	//var c *int = &b
	//fmt.Printf("%s", p)
	//fmt.Println(p)
}
