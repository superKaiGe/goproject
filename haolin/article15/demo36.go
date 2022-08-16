package main

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name: name}
}
func (dog *Dog) SetName(name string) {
	dog.name = name
}
func main() {
	//示例1.
	//New("little pig").SetName("lucky") //不能调用不可寻址的值的指针方法
	//fmt.Printf("%s", dog.name)
	//示例2.
	//a:= map[string]int{"the": 0}["the"]++
	//fmt.Printf("%s",a)
	//map1 := map[string]int{"the": 0}
	//fmt.Printf("%s",map["the"]++)
	//fmt.Printf("%s",map1)
}
