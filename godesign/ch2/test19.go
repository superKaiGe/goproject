package main

import "fmt"

func main() {
	//s := []int{}
	var s []int
	//s = nil
	//s = []int(nil)
	s = []int{}
	if s == nil {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}
	slice1 := make([]int, 1, 3)
	if slice1 == nil {
		fmt.Println("ccc")
	} else {
		fmt.Println("ddd")
	}
	//fmt.Println("%T", slice1)
}
