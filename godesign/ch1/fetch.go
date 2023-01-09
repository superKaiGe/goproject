package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//for _, url := range os.Args[1:] {
	//	resp, err := http.Get(url)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
	//		os.Exit(1)
	//	}
	//	//b, err := ioutil.ReadAll(resp.Body)
	//	defer resp.Body.Close()
	//	//if err != nil {
	//	//	fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
	//	//	os.Exit(1)
	//	//}
	//	//var a io.Writer
	//	data, err := io.Copy(os.Stderr, resp.Body)
	//	if err != nil {
	//		fmt.Printf("err:%s", err)
	//	}
	//	fmt.Println("write", data)
	//	fmt.Println(resp.Status)
	//fmt.Printf("%s", result)
	//return
	//fmt.Printf("%s", b)
	//}
	url := "gimg3.baidu.com/search/src=http%3A%2F%2Fpics7.baidu.com%2Ffeed%2F9345d688d43f879468fd56957ac83fff1ad53a6d.jpeg%40f_auto%3Ftoken%3D9f6c9e4ebb91eb332ac3059d686b4442&refer=http%3A%2F%2Fwww.baidu.com&app=2021&size=f360,240&n=0&g=0n&q=75&fmt=auto?sec=1667494800&t=889dc7be5e71ccae4d54897fd6ceeb5b"
	resp, err := http.Get(url)
	res := strings.HasPrefix(url, "https1://")
	fmt.Println(url)
	if !res {
		url = "https://" + url
	}
	fmt.Println(url)
	return
	fmt.Printf("%v", resp)
	return
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()

	out, err := os.Create("2.png")
	wt := bufio.NewWriter(out)

	defer out.Close()

	n, err := io.Copy(out, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}
	wt.Flush()
	//fmt.Printf("%s", os.Stderr)
	//}
}
