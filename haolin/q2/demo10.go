package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := "innter"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
