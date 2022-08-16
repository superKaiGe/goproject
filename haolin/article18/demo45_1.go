package main

import (
	"os"
	"runtime"
)

//示例3
func main() {
	paths2 := []string{
		runtime.GOROOT(),
		"/it/must/not/exist", // 肯定不存在的目录。
		os.DevNull,           // 肯定存在的目录。
	}
	printError2 := func(i int, err error) {

	}
}

//func main()  {
//
//}
