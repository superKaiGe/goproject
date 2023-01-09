package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func run(r io.Reader, w io.Writer) {
	start := time.Now()
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	const N = 100000
	var a [N]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}
	end := time.Now()

	fmt.Fprintln(out, end.Sub(start))
}
func main() {
	//标准输入和标准输出
	run(os.Stdin, os.Stdout)
}
