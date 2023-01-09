package main

import "fmt"

func main() {
	var intChan = make(chan int, 10)
	var strChan = make(chan string, 10)
	select {
	case e1 := <-intChan:
		fmt.Printf("The 1th case was selected.e1=%v\n", e1)
	case e2 := <-strChan:
		fmt.Printf("The 2nd case was selected.e2=%v\n", e2)
	default:
		fmt.Println("Default!")
	}
}
