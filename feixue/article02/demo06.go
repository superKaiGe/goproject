package main

import "fmt"

func main() {
	ages := map[string]int{"张三": 20}
	fmt.Println(ages)
	modify(ages)
	fmt.Println(ages)
}
func modify(m map[string]int) {
	m["张三"] = 30
}
