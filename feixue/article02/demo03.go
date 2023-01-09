package main

import "fmt"

func main() {
	name := "张三"
	fmt.Println(modify(name))
	fmt.Println(name)
}

func modify(s string) string {
	s = s + s
	return s
}
