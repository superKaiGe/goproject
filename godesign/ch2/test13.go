package main

import "fmt"

func main() {
	var a [3]int             //3个整数的数组
	fmt.Println(a[0])        //输出数组的第一个元素
	fmt.Println(a[len(a)-1]) //输出数组的最后一个元素,即是[2]
	// 输出索引和元素
	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
	//仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	fmt.Println(q[2])
	//q = [...]int{1, 2, 3, 4}
	//q = [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", q)
	fmt.Println(".................")
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "$1", GBP: "$2", RMB: "$3"}
	fmt.Println(RMB, symbol[RMB])
}
