package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "hello,everyone!\n") //这里对err进行了重声明
	if err != nil {
		fmt.Printf("Error:%v\n", err)
	}
	fmt.Printf("%d byte(s) were writern.\n", n)
}
