package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("你好"))
	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b, ",", "http://www.flysnow.org")
	//把Buffer里的内容打印到终端控制台
	b.WriteTo(os.Stdout)
	//println(b)
	//var p [100]byte
	//n, err := b.Read(p[:])
	//fmt.Println(n, err, string(p[:n]))
	println()
	//var p [100]byte
	//p[0] = 1
	//n, err := b.Read(p[:])
	//fmt.Println(n, err, string(p[:n]))

	data, err := ioutil.ReadAll(&b)
	fmt.Println(data, err)
}
