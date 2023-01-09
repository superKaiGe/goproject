package main

import "fmt"

func main() {
	array := [5]int{1: 2, 3: 4}
	modify(&array)
	fmt.Println(array)
}
func modify(a *[5]int) {
	a[1] = 3
	fmt.Println(*a)
}
