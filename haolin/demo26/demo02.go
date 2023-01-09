package main

import (
	"fmt"
	"time"
)

func main() {
	//name := "Eric"
	//go func() {
	//	fmt.Printf("Hello, %s.\n", name)
	//}()
	//runtime.Gosched()
	//name = "Harry"
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s.\n", name)
		}()
		time.Sleep(time.Millisecond)
	}
	//runtime.Gosched()

}
