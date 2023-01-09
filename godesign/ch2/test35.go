package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapv := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {

		//str :=
		if input.Text() == "" {
			break
		}
		mapv[input.Text()]++

	}

	fmt.Println("The length %d", len(mapv))
	for k, v := range mapv {
		fmt.Printf("%s %d\n", k, v)
	}
}
