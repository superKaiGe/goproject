package main

import (
	"fmt"
	"strconv"
)

func main() {
	//整数变为字符串
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 10))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
	//字符串变为整数
	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(x)
	y1, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(y1)

}
