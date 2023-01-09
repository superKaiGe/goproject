package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}
func add(a, b int) int {
	return a + b
}
