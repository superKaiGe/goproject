package main

import "fmt"

func main() {
	outerFunc()
}
func outerFunc() {
	defer fmt.Println("函数执行结束前一刻才会被打印")
	fmt.Println("第一个被打印.")
}
