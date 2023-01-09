package cow

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// ConcurrentArray 代表并发安全的整数数组接口。
type ConcurrentArray interface {
	// Set 用于设置指定索引上的元素值。
	Set(index uint32, elem int) (err error)
	// Get 用于获取指定索引上的元素值。
	Get(index uint32) (elem int, err error)
	//Len 获取数组的长度
	Len() uint32
}
type intArray struct {
	length uint32
	val    atomic.Value
}

func (array *intArray) Set(index uint32, elem int) (err error) {
	//TODO implement me
	//panic("implement me")
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	newArray := make([]int, array.length)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *intArray) Get(index uint32) (elem int, err error) {
	//TODO implement me
	//panic("implement me")
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *intArray) Len() uint32 {
	//TODO implement me
	//panic("implement me")
	return array.length
}

func (array *intArray) checkIndex(index uint32) error {
	if index >= array.length {
		return fmt.Errorf("Index out of range [0, %d)!", array.length)
	}
	return nil
}

// checkValue 用于检查原子值中是否已存有值。
func (array *intArray) checkValue() error {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}
func NewConcurrentArray(length uint32) ConcurrentArray {
	array := intArray{}
	array.length = length
	array.val.Store(make([]int, length))
	return &array
}
