package main

import "fmt"

type person struct {
	name string
}

func (p person) string() string {
	return "the person name is " + p.name
}
func main() {
	p := person{name: "张三"}
	fmt.Println(p.string())
	(&p).modify()
	fmt.Println(p.string())
}
func (p *person) modify() {
	p.name = "李四"
}
