package main

import "fmt"

func main() {
	//示例1
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println()
	//示例2
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("the slice:%v\n", slice1)
	slice2 := modifyslice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2)
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println()

	//modifyComplexArray
	//示例3
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The complex array: %v\n", complexArray1)
	fmt.Printf("The complex array: %v\n", complexArray2)
}
func modifyArray(a [3]string) [3]string {
	a[1] = "X"
	return a
}
func modifyslice(a []string) []string {
	a[1] = "i"
	return a
}
func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	return a
}
