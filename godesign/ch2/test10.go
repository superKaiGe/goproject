package main

import "fmt"

func main() {
	var a [3]int             //3个整数的数组
	fmt.Println(a[0])        //输出数组的第一个元素
	fmt.Println(a[len(a)-1]) //输出数组的最后一个元素
	//输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	fmt.Println()
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Printf("%T", q)
	fmt.Println()
	fmt.Printf("%v", q)
	fmt.Println()
	fmt.Println(r[2])
	q1 := [3]int{1, 2, 3}
	fmt.Println(q1)
	//q1 = [4]int{1, 2, 3, 4}
	type Currency int
	const (
		USD Currency = iota
		EUR
	)
	symbol := [...]string{USD: "$", EUR: "$1"}
	fmt.Println(EUR, symbol[EUR])
	r1 := [...]int{99: -1}
	fmt.Println(r1)
}
