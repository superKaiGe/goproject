package main

import "fmt"

func main() {
	chs := make(chan string, 2)
	chs <- "first"
	chs <- "second"
	ch_len := (len(chs))
	for i := 0; i < ch_len; i++ {
		fmt.Println(<-chs)
	}
	//for i := range chs {
	//	fmt.Println(i)
	//}
}
