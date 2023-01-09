package main

import "fmt"

const (
	AiB = 2 << (10 * iota)
	KiB
	MiB
	YiB
)

func main() {
	//fmt.Println(YiB / ZiB)
	fmt.Println(AiB, KiB, MiB, YiB)
}
