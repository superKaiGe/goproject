package main

func main() {
	var a int = 5
	var p1 *int = &a //p1指向变量a所在的内存单元
	var p2 *int = &a //p2指向变量a所在的内存单元
	(*p1) += 6       //通过p1修改变量a的值
	println(*p1)
	println(*p2) //对p1的修改可以通过另一个指针变量p2的解引用反映出来
}
