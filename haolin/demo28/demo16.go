package main

import "fmt"

func main() {
	var strChan = make(chan string, 10)
	//strChan <- "a"
	select {
	case e, ok := <-strChan:
		if !ok {
			fmt.Println("End.")
			break
		}
		fmt.Printf("Received:%v\n", e)
	}
}
