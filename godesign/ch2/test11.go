package main

import "fmt"

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
func test1(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func main() {
	primes := []int{1, 3, 5, 7}
	primes_1 := []int{9, 11}
	primes = append(primes, primes_1...)
	primes = append(primes, 1)
	fmt.Printf("%T\n", primes)
	fmt.Println(Sum(primes...))
	fmt.Println()
	strss := []string{"a", "b", "c"}
	strss1 := []string{"d", "e", "f"}
	strss = append(strss, strss1...)
	test1(strss...)
	fmt.Println()
	//var q [3]int = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	fmt.Println(r[2])
	q := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", q)
}
