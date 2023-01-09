package main

import (
	"fmt"
	"time"
)

//const pi = 3.14159
//
//const (
//	e  = 2.718281
//	pi = 3.1415
//)
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
const (
	a = 1
	b
	c = 2
	d
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp        Flags = 5 << iota //向上
	FlagBroadcast                   //支持广播访问
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println(a, b, c, d)
	fmt.Println(Sunday, Monday, Tuesday)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}
