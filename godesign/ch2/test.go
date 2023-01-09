package main

import "fmt"

func main() {
	x := "hello"
	//fmt.Println(len(x))
	//return
	//for i := 0; i < len(x); i++ {
	//	x := x[i]
	//	//fmt.Println(x)
	//	//fmt.Println('!')
	//	if x != '!' {
	//		x := x + 'A' - 'a'
	//		fmt.Printf("%c", x)
	//		//fmt.Println(x)
	//	}
	//}
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println()
}
