package common

import "fmt"

type count int

func New1(v int) count {
	return count(v)
}

type Loginer interface {
	Login()
}
type defaultLogin int

func (d defaultLogin) Login() {
	fmt.Println("login in...")
}
func NewLoginer() Loginer {
	return defaultLogin(0)
}
