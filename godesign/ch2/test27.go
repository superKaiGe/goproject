package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice": 30,
	}
	age, ok := ages["bob"]
	if !ok {
		age = 111
	}
	fmt.Println(age)
	RES := equal(map[string]int{"A": 0}, map[string]int{"A": 0})
	fmt.Println(RES)
}
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
