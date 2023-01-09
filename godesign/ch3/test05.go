package main

import "fmt"

func main() {
	//s := []int{}
	//第二个参数是slice的len
	//var s = make([]int, 1)
	//s[0] = 1
	//fmt.Println(s)
	var s []int //nil
	//if s == nil {
	//	fmt.Println("this is a nil slice")
	//}
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

}
