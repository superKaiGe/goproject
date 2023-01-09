package main

import "fmt"

func main() {
	//m := map[string]string{} 字面量的方式
	//var m = map[string]string{} 字面量的方式
	m := make(map[string]string)
	//var m map[string]string //nil map
	m["name"] = "laughing"
	fmt.Println(m)

	if m == nil {
		fmt.Println("this map is a nil map")
	}
}
