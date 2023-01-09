package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d\n", i)
		}()
	}
	printNUmbers()
}
func printNUmbers() {
	for i := 0; i < 6; i++ {
		defer func(n int) {
			fmt.Printf("%d\n", n)
		}(i)
	}
}
