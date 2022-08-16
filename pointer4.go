package main

func main() {
	var a int = 5
	var b int = 6
	var p *int = &a //指向变量a所在内存单元
	println(*p)     //输出变量a的值
	p = &b          //指向变量b所在内存单元
	println(*p)     //输出变量b的值
}
