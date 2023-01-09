package main

import "fmt"

func main() {
	s := "hello, world你"
	fmt.Println(len(s)) //12
	//fmt.Println(s[0], s[8], s[12]) //104 119
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("googbye" + s[5:])

	s1 := "left foot"
	t := s1
	s1 += ", right foot"
	fmt.Println(s1)
	fmt.Println(t)
	//s1[0] = 'L'
}
