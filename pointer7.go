package main

func main() {
	var a int = 5
	var p *int = &a
	println(*p)
	foo1(&p)
	println(*p)
}
func foo1(pp **int) {
	var b int = 55
	var p1 *int = &b
	*pp = p1
}
