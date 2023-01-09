package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[2])
	slice[2] = 10 //修改值
	fmt.Println(slice[2])
	//fmt.Println(slice[-1])
	slice = []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 10, 20)
	fmt.Println(newSlice)
	fmt.Println(slice)
	fmt.Println("++++++++++")
	slice = []int{1, 2, 3, 4, 5}
	newSlice = slice[1:2:3]
	newSlice = append(newSlice, slice...)
	fmt.Println(newSlice)
	fmt.Println(slice)

}
