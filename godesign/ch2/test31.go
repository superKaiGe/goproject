package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const N = 10000
	var a [N]int
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
