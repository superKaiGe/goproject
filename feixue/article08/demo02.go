package main

import (
	"goproject/feixue/common"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutine = 5
	pooRes       = 2
)

func main() {
	//等待任务完成
	//var wg sync.Mutex
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p, err := common.New2(createConnection, pooRes)
	if err != nil {
		log.Println(err)
		return
	}
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}
func dbQuery(query int, pool *common.Pool) {
	coon, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(coon)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, coon.(*dbConnection).ID)
}

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{
		id,
	}, nil
}
