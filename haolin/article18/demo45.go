package main

import (
	"fmt"
	"os"
	"os/exec"
)

//underlyingError 会返回已知的操作系统相关错误的潜在错误值
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}
func main() {
	//示例1
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n",
		uError, uError)
	fmt.Println()
	//示例2
	paths := []string{
		os.Args[0],           //当前源码文件或可执行文件
		"/it/must/not/exist", //肯定不存在的目录
		os.DevNull,           //肯定存在的目录
	}
	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Printf("error(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		}
	}
	var f *os.File
	var index int

	index = 0
	f, err = os.Open(paths[index])
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	f.Close()
	_, err = f.Read([]byte{})
	printError(index, err)
	if f != nil {
		f.Close()
	}
	fmt.Println()

}
