package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
}
func caller() {
	fmt.Println("Enter function caller")
	panic(errors.New("something wrong"))
	//panic(fmt.Println)
	fmt.Println("Exit function caller")
}
