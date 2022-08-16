package main

import (
	"flag"
	"fmt"
)

//func main() {
//	var name string
//	flag.StringVar(&name, "name", "everyone", "The greeting object.")
//	flag.Parse()
//	fmt.Printf("Hello,%s\n", name)
//}
func main() {
	//var name = flag.String("name", "everyone", "the greeting object.")
	//类型断言 不给变量指定类型 让程序自己判断
	//name := flag.String("name", "everyone", "the greeting object.") //类型推断,短变量声明
	//name1 := "sss"                                                  //短变量声明 只能在函数体内部使用
	//flag.Parse()
	//fmt.Printf("hello,%v\n", *name)
	//fmt.Printf("hello,%s\n", *name)
	//fmt.Printf("%s\n", name1)
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("%s\n", *name)
}
func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

//func main() {
//
//}
