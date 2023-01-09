package main

import "fmt"

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // slice[2] = 9
	// 5 6 9 8 9
	return slice[:len(slice)-1]
}
func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	s1 := []int{5, 6, 7, 8, 10}
	fmt.Println(remove1(s1, 2))
}
func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice[i:])
	return slice[:len(slice)-1]
}
