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
type Hello interface {
	hello()
}

func (u user) hello() {
	fmt.Println("Hello, i am a user")
}
func sayHello(h Hello) {
	h.hello()
}
func main() {
	ad := admin{user{"skg", "skg_email"}, "主管"}
	sayHello(ad.user)
	sayHello(ad)
}
