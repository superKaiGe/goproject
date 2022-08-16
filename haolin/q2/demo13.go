package main

import "fmt"

func main() {
	//a := string(-1)
	a := string(-1)
	fmt.Println(a)
	b := string([]rune{'\u4F60', '\u597D'})
	fmt.Println(b)
	var c byte
	fmt.Println(c)
	//Mystring 是string的别名类型 别名类型与源类型的区别只在名称上 他们是完全相同的,
	//别名类型主要是为了代码重构而存在的 byte是unit8的别名类型 run是int32的别名类型
	type Mystring = string
	//Mystring1 和string 是不一样的类型 新类型,不同于其他类型 类型的再定义
	type Mystring1 string
	var e string
	var f Mystring1
	g := Mystring1(e)
	if f == g {
		println("OK")
	}
}
