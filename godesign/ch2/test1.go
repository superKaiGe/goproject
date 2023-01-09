package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed:%v", err)
	}
	//log.Printf("Working directory = %s", cwd)
	//fmt.Printf("working directory1 = %s", cwd)
	//fmt.Println(cwd)
}
func main() {

}
