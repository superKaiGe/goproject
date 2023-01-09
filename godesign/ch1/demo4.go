package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//fmt.Println("aaa")
	//fmt.Println(os.Args[0])
	fmt.Println(time.Now())
	fmt.Println(strings.Join(os.Args[1:], " d "))
	fmt.Println(time.Now())
}
