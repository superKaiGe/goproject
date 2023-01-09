package main

import "fmt"

func main() {
	a := gcd(2, 5)
	fmt.Println(a)
	b := fib(5)
	fmt.Println(b)
}
func gcd(x, y int) int {
	for y != 0 {
		//2 5 5 2
		//5 2 2 1
		//2 1 1 0
		x, y = y, x%y
		//5 2
	}
	return x
}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Printf("%v-%v\n", x, y)
	}
	return x
}
