package q3

import (
	"flag"
)

var name string

func init() {
	//flag.StringVar(&name, "name", "everyone", "The greeting project")
}
func Main_demo5() {
	flag.Parse()
	//fmt.Printf("name %s\n", name)
	hello(name)
}
