package main

import (
	"fmt"
	"math"
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	const Pi64 float64 = math.Pi
	var x1 float32 = float32(Pi64)
	var y1 float64 = Pi64
	var z1 complex128 = complex128(Pi64)
	fmt.Println(x1, y1, z1)
	// 0 0.0 0i '\u0000'
	fmt.Println()
	var f float64 = 212
	//fmt.Printf("%T", (f-32)*5/9)
	fmt.Println(5 / 9 * (f - 32))
	fmt.Printf("%T", 5/9*(f-32))
	fmt.Println()
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}
