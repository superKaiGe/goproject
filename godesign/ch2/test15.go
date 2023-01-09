package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	ptr := [32]byte{1}
	fmt.Println(ptr)
	zero1(&ptr)
	fmt.Println(ptr)
	zero2(&ptr)
	fmt.Println(ptr)
}
func zero1(ptr *[32]byte) {
	*ptr = [32]byte{}
}
func zero2(ptr *[32]byte) {
	*ptr = [32]byte{2, 1}
}
func zero3(ptr *[32]byte) {

}
