package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " d "
	}

	fmt.Println(s)
	fmt.Println(time.Now())
}
