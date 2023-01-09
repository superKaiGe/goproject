package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"os"
)

func main() {
	file, err := os.Open("/notexists")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file)
}
func add(a, b int) (int, error) {
	return a + b, nil
}
