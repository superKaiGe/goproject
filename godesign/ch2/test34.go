package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//分行读取标准输入s，将s放入map统计出现次数
	input := bufio.NewScanner(os.Stdin)
	statistic := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		statistic[line]++
		fmt.Println(statistic)
	}
	//打印出现重复的输入
	for line, num := range statistic {
		//if num > 1 {
		fmt.Printf("%s\t%d\n", line, num)
		//}
	}
}
