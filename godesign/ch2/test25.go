package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	//delete(ages, "alice") 删除元素alice
	fmt.Println(ages["alice"]) //map元素访问也是通过下标的方式
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"]) //map元素访问也是通过下标的方式
	ages["bobb"] += 1
	ages["bobbc"]++
	fmt.Println(ages["bobb"])
	fmt.Println(ages["bobbc"])
	fmt.Println(&ages)
	fmt.Println("--------------")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
