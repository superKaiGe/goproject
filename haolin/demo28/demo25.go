package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		fmt.Println("try to recover the panic.")
		if p := recover(); p != nil {
			fmt.Printf("Recovered the panic(%#v).\n", p)
		}
	}()
	var mutex sync.Mutex
	fmt.Println("Lock the lock.")
	mutex.Lock()
	fmt.Println("The lock is locked.")
	fmt.Println("unlock the lock")
	mutex.Unlock()
	fmt.Println("the lock is unlock.")
	fmt.Println("unlock the lock again.")
	mutex.Unlock()
}
