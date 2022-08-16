package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	//container := map[string]int{0: "zero"}
	container := map[int]string{0: "zero", 1: "oneoeone", 2: "two"}
	value, _ := interface{}(container).([]string)
	if value != nil {
		fmt.Printf("The element is %q.\n", value)
	}
}
