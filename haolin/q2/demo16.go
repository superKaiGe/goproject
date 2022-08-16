package main

import "fmt"

func main() {
	//          0  1  2  3  4  5  6  7
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:7]
	s5 := s4[0:cap(s4)]
	fmt.Printf("%d\n", len(s3))
	fmt.Printf("%d\n", cap(s3))
	fmt.Printf("%d\n", s3)
	fmt.Printf("%d\n", len(s4))
	fmt.Printf("%d\n", cap(s4))
	fmt.Printf("%d\n", s4)

	fmt.Println("----------------")
	fmt.Printf("%d\n", len(s5))
	fmt.Printf("%d\n", cap(s5))
	fmt.Printf("%d\n", s5)
}
