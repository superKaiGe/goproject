package main

import "fmt"

func main() {
	chs := make(chan string, 2)
	chs <- "first"
	chs <- "second"
	for i2 := range chs {
		fmt.Println(i2)
		//fmt.Println(i)
	}
}
