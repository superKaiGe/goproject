package module

import (
	"math"
	"sync"
)

// SNGenertor 代表序列号生成器的接口类型。
type SNGenertor interface {
	// Start 用于获取预设的最小序列号。
	Start() uint64
	// Max 用于获取预设的最大序列号。
	Max() uint64
	// Next 用于获取下一个序列号。
	Next() uint64
	// CycleCount 用于获取循环计数。
	CycleCount()
	// Get 用于获得一个序列号并准备下一个序列号。
	Get() uint64
}

// NewSNGenertor 会创建一个序列号生成器。
// 参数start用于指定第一个序列号的值。
// 参数max用于指定序列号的最大值。
func NewSNGenertor(start uint64, max uint64) SNGenertor {
	if max == 0 {
		max = math.MaxUint64
	}
	return &mySNGenertor{
		start: start,
		max:   max,
		next:  start,
		//cycleCount: 0,
		//lock:       sync.RWMutex{},
	}
}

// mySNGenertor 代表序列号生成器的实现类型。
type mySNGenertor struct {
	// start 代表序列号的最小值。
	start uint64
	// max 代表序列号的最大值。
	max uint64
	// next 代表下一个序列号。
	next uint64
	// cycleCount 代表循环的计数。
	cycleCount uint64
	// lock 代表读写锁。
	lock sync.RWMutex
}

func (m mySNGenertor) Start() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m mySNGenertor) Max() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m mySNGenertor) Next() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m mySNGenertor) CycleCount() {
	//TODO implement me
	panic("implement me")
}

func (m mySNGenertor) Get() uint64 {
	//TODO implement me
	panic("implement me")
}
