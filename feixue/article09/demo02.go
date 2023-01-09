package main

import "sync"

type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

//存储操作
func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

//获取操作
func (sm SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

//删除操作
func (sm SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

//变量map 并且把变量的值给回调函数 可以让调用者控制做任何事情
func (sm SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLocker()
	defer sm.rw.RUnlock()
	for k, v := range sm.data {
		cb(k, v)
	}
}

//生成一个初始化的SynchronizedMap
func newSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}
