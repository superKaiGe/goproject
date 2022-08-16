package main

import "fmt"

func main() {
	//示例1
	//numbers1 := []int{1, 2, 3, 4, 5, 6}
	//for i := range numbers1 {
	//	if i == 3 {
	//		numbers1[i] |= i
	//	}
	//}
	//fmt.Println(numbers1)
	//fmt.Println()
	//示例2
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := numbers2
	fmt.Println(numbers3)
	maxIndex2 := len(numbers3) - 1
	// 0+1 =>1 0的值+1的值 7 3 5 7 9 11
	// 1 3 3 4 5 6
	// 1 3 6 4 5 6
	// 1 3 6 10 5 6
	// 1 3 6 10 15 6
	// 1 3 6 10 15 21
	// 22 3 6 10 15 21
	fmt.Println(maxIndex2)
	for i, e := range numbers3 {

		if i == maxIndex2 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println(numbers3)
	fmt.Println()
}
