package main

import "fmt"

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
	}()
	return <-responses
}

func request(hostname string) (response string) {
	return hostname
}
func main() {
	hostname := mirroredQuery()
	fmt.Println(hostname)
}
