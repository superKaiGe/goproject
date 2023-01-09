package main

import "fmt"

type user struct {
	name  string
	email string
}
type admin struct {
	user
	level string
}

func (u user) sayHello() {
	fmt.Println("Hello,i am a user")
}
func (a admin) sayHello() {
	fmt.Println("Hello,i am a admin")
}
func main() {
	ad := admin{
		user{"张三", "zhangsan@flysnow.org"},
		"管理员",
	}
	ad.user.sayHello()
	ad.sayHello()
}

//从输出看
