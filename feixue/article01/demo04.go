package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Printf("索引:%d,值:%d\n", i, v)
	}
	fmt.Println("-----")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("索引:%d,值:%d\n", i, slice[i])
	}
}
