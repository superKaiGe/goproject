package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	jim := person{
		age:  10,
		name: "Jim",
	}
	fmt.Println(jim)
	modify(&jim)
	fmt.Println(jim)

}
func modify(p *person) {
	p.age = p.age + 10
}
