package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main")
	//引发panic
	panic(errors.New("something wrong"))
	p := recover()
	fmt.Printf("panic,%s\n", p)
	fmt.Println("Exit function main.")
}
