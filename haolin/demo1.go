package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
}

func main() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("hello,%s!\n", name)
}
