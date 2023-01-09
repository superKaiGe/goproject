package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//准备从标准设备读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	//读取数据直到碰到\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred:%s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("Hello,%s! What can I do for you?\n", name)
	}

	for true {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye")
			os.Exit(1)
		default:
			fmt.Println("Sorry,I didn't catch you.")

		}

	}

}
