package main

import "fmt"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func main() {
	a := []string{"1"}
	b := []string{"1", "2"}
	c := equal(a, b)
	if c {
		fmt.Println("success")
	} else {
		fmt.Println("error")
	}
}
