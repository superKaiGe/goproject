package lib

import (
	"errors"
	"fmt"
)

//Gotickets 表示 goroutine票池的接口
type GoTickets interface {
	//拿走一张票
	Take()
	//归还一张票
	Return()
	//票池是否已被激活
	Active() bool
	//票的总数
	Total() uint32
	//剩余的票数
	Remainder() uint32
}

//myGotickets 表示goroutine票池的实现
type myGoTickets struct {
	total    uint32        //票的总数
	ticketCh chan struct{} //票的容器
	active   bool          //票池是否已经被激活
}

func (gt *myGoTickets) Take() {
	//TODO implement me
	//panic("implement me")
	<-gt.ticketCh
}

func (gt *myGoTickets) Return() {
	//TODO implement me
	//panic("implement me")
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	//TODO implement me
	//panic("implement me")
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	//TODO implement me
	//panic("implement me")
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	//TODO implement me
	//panic("implement me")
	return uint32(len(gt.ticketCh))
}

func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg :=
			fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}
func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}
