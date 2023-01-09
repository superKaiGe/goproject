package main

import (
	"os"
	"strconv"
)

func main() {
	//create 函数创建一个不存在的文件
	file, _ := os.Create("data.txt")
	//函数返回的时候关闭文件句柄
	defer file.Close()
	for i := 0; i < 100000; i++ {
		file.WriteString(strconv.Itoa(i%10) + "\n")
	}
}
