package main

import (
	"goproject/feixue/common"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("...开始执行任务...")
	timeout := 4 * time.Second
	r := common.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case common.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case common.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}
func createTask() func(int) {
	return func(i int) {
		log.Printf("正在执行任务%d", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
