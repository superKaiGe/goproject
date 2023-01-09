package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d\n", n)
		}(i * 2)
	}
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d\n", n)
		}(i * 2)
	}
}
