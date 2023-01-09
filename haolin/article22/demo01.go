package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 0, "It indicates whether to use a mutex to protect data writing")
}
func main() {
	flag.Parse()
	//buffer 代表缓冲区
	var buffer bytes.Buffer
	const (
		max1 = 5
		max2 = 10
		max3 = 10
	)
	//mu 代表以下流程要使用的互斥锁
	var mu sync.Mutex
	//sign 代表信号的通道
	sign := make(chan struct{}, max1)
	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= max2; j++ {
				//准备数据
				header := fmt.Sprintf("\n[id: %d, iteration: %d]",
					id, j)
				data := fmt.Sprintf(" %d", id*j)
				//写入数据
				if protecting > 0 {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}
	for i := 0; i < max1; i++ {
		<-sign
	}
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error:%s", err)
	}
	log.Printf("The contents:\n%s", data)
}
