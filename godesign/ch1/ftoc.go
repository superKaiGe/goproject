package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF=%C\n", freezingF, fToc(freezingF)) // 32F =0C
	fmt.Printf("%gF=%C\n", boilingF, fToc(boilingF))   // 32F =0C

}
func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
