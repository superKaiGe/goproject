package main

import "fmt"

func main() {
	//var m map[string]int
	m := map[string]int{
		string(1): 2,
	}
	//m:=
	key := "two"
	m["two"] = 2
	elm, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map:%d(%v)\n", key, elm, ok)
	//fmt.Printf("The length of nil map;%d\n", len(m))
	//fmt.Printf("Delete the key-element pair by key %q...\n", key)
	//delete(m, key)
	//elm = 2
	//fmt.Println("Add a key-element pair to a nil map...")
	//m["two"] = elm
}
