package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Hello World")
	var w io.Writer
	w = &b
	fmt.Println(w)
}
