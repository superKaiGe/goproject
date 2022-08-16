package main

import (
	"flag"
	"fmt"
	"os"
)

var hostname string
var username string
var pwd string
var port string

func init() {
	flag.StringVar(&hostname, "hostname", "127.0.0.1", "server ip or hostname")
	flag.StringVar(&username, "username", "root", "the user name to login database")
	flag.StringVar(&pwd, "pwd", "", "the password of username")
	flag.StringVar(&port, "port", "3306", "the port of your mysql")
}
func main() {

	//flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s\n", "Link to your Mysql")
	}
	fmt.Printf("hostname: %s\nusername: %s\npassword:  %s\nport:  %s\n", hostname, username, pwd, port)
}
