package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

type Point struct {
	X, Y, C int
}

//就地排序
func sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	data := appendValues(values[:0], root)
	return data
}

//appendValues 将元素按照顺序追到到 values 里面,然后返回结果slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func main() {

	values := sort([]int{1, 3, 2, 5, 4})
	fmt.Println(values)
	p := Point{
		X: 1,
		Y: 2,
		C: 3,
	}
	fmt.Println(p)
}
