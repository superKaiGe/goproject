package main

import "fmt"

func main() {
	//指定切片的长度 这个时候切片的容量也是5
	_ = make([]int, 5)
	//使用内置的make函数 指定切片的长度 容量是10对应的是切片底层数组的
	_ = make([]int, 5, 10)
	//使用字面量 指定初始化的值
	_ = []int{1, 2, 3, 4}
	//只初始化某个索引的值
	_ = []int{4: 1}

	array1 := [5]int{4: 1}
	slice1 := []int{4: 1}
	fmt.Println(array1)
	fmt.Println(slice1)
	//nil切片
	var nilSlice []int
	//空切片
	slice2 := []int{}
	fmt.Println(&nilSlice)
	fmt.Println(&slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)
	slice4 := slice3[:]
	slice5 := slice3[0:]
	slice6 := slice3[:5]
	slice6[0] = 11
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := slice7[1:3]
	fmt.Printf("slice8长度:%d,容量:%d", len(slice8), cap(slice8))
	fmt.Println("-----------分割线-------")
	slice9 := []int{1, 2, 3, 4}
	slice10 := slice9[1:2:3]
	fmt.Printf("slice10长度:%d,容量:%d,", len(slice10), cap(slice10))
}
