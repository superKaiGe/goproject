package main

//
//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//	runtime.GOMAXPROCS(1)
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for i := 1; i < 1000; i++ {
//			fmt.Println("A:", i)
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 500; i++ {
//			fmt.Println("B:", i)
//		}
//	}()
//	wg.Wait()
//}
