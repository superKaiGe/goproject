package main

import "fmt"

//迭代器
func iterator(iterable []interface{}) chan interface{} {
	yield := make(chan interface{})
	go func() {
		for i := 0; i < len(iterable); i++ {
			yield <- iterable[i]
		}
		close(yield)
	}()
	return yield
}

//获取下一个元素
func next(iter chan interface{}) interface{} {
	for v := range iter {
		return v
	}
	return nil
}
func main() {
	nums := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
	iter := iterator(nums)
	for v := next(iter); v != nil; v = next(iter) {
		fmt.Println(v)
	}
}
