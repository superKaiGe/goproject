package main

import "time"

var a string

func f() {
	print(a)
}

//func hello() {
//
//}
func main() {
	a = "hello,world"

	go f()
	time.Sleep(time.Second)
	//hello()
	//time.Sleep(time.Second * 2)
}
