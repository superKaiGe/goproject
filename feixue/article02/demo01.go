package main

import (
	"fmt"
	"sort"
)

func main() {
	dict := make(map[string]int)
	dict["张三"] = 43
	dict = map[string]int{"张三": 29}
	dict = map[string]int{"张三": 29, "李四": 30}
	dict = map[string]int{}
	var dict1 map[string]int //nil map
	dict1 = make(map[string]int)
	dict1["张三"] = 1
	delete(dict1, "张三")
	delete(dict1, "李四")
	fmt.Println(dict1["张三"])
	age, exits := dict["李四"]
	if exits {
		fmt.Println(age)
	} else {
		fmt.Println("key 李四不存在")
	}
	dict = map[string]int{"张三": 27, "王五": 28, "李四": 29, "赵六": 30}
	for s, i := range dict {
		fmt.Println(s, i)
	}
	fmt.Println("----------------星期五-----------------")
	var names []string
	for name := range dict {
		names = append(names, name)
	}
	sort.Strings(names) //排序
	for _, key := range names {
		fmt.Println(key, dict[key])
	}

}
