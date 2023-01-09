package main

func main() {
	ch := make(chan string, 2)
	ch <- "first"
	ch <- "second"
	close(ch)
	ch <- "3"
}
