package main

import "fmt"

func main() {
	//outerFunc()
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovered panic:%s\n", p)
		}
	}()
	myIndex := 3
	ia := [3]int{1, 2, 3}
	_ = ia[myIndex]
	fmt.Printf("Recovered panic:%s\n", 2)

	//outerFunc()
}

//func outerFunc() {
//	interFunc()
//}
//func interFunc() {
//	panic(errors.New("An intended fatal error"))
//}
