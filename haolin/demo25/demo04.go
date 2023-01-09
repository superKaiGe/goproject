package main

var done = make(chan bool, 1)
var msg string

func aGoroutine() {
	msg = "hello, world"
	<-done
}

func main() {
	go aGoroutine()
	done <- true
	println(msg)
}
