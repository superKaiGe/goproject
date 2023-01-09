package main

type person struct {
	age  int
	name string
}

var p person

func main() {
	_ := person{10, "jim"}
	_ := person{
		name: "jim",
		age:  10,
	}

}
