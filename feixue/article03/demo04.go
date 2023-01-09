package main

import "fmt"

func main() {
	fmt.Println("1", "2", "3")
	print("1", "2", '3')
}
func print(a ...interface{}) {
	for _, i2 := range a {
		fmt.Print(i2)
	}
	fmt.Println()
}
