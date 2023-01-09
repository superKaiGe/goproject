package main

import (
	"fmt"
	"goproject/feixue/common"
)

func main() {
	//c := common.Count(1)
	c := common.New(100)
	fmt.Println(c)
	login := common.NewLoginer()
	login.Login()
}
