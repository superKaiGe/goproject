package main

import "fmt"

func generate(ch chan int, n int) {
	for i := 2; i <= n; i++ {
		fmt.Println("generate:", i)
		ch <- i
	}
	close(ch)
}
func filter(in, out chan int, prime int) {
	for i := range in {
		fmt.Printf("filter(%d): %d\n", prime, i)
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

func main() {
	res := []int{}
	ch := make(chan int)
	go generate(ch, 112)
	for {
		if prime, ok := <-ch; ok {
			res = append(res, prime)
			ch_out := make(chan int)
			go filter(ch, ch_out, prime)
		} else {
			break
		}
	}
	fmt.Println("main:", res)
}
