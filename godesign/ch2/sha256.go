package main

import (
	"fmt"
)

func main() {
	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Printf("%T\n%T\n%v\n%v\n", []byte("aaa"), []byte{1}, []byte("aaa"), []byte{byte(120)})
	fmt.Println([]byte("aaa"))
	c3 := [32]byte{1}
	fmt.Println(c3)
	zero(&c3)
	fmt.Println(c3)
}
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 2
	}
}
