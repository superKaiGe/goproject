package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question111")
		flag.PrintDefaults()
	}
}
func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	var cmdline = flag.NewFlagSet("question", flag.ExitOnError)
	cmdline.StringVar(&name, "name", "everyone", "sss")
	cmdline.Parse(os.Args[1:])
	//cmdline.Var('111', "skg", "sss")
	//flag.StringVar(&name, "name", "everyone", "sss")
	//flag.Parse()
	fmt.Printf("hello,%s\n", name)
}
