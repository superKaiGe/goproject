package main

import "fmt"

func main() {
	var intChan chan int
	var strChan chan string

	select {
	case e1 := <-intChan:
		fmt.Printf("e1=%v\n", e1)
	case e2 := <-strChan:
		fmt.Printf("e2=%v\n", e2)
	default:
		//fmt.Printf("%v\n","default")
		fmt.Println("default")
	}
}
