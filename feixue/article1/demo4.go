package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

var key = "name"

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancle()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
