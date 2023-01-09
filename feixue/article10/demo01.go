package main

import "log"

func init() {
	log.SetPrefix("【UserCenter】")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
func main() {
	log.Println("这是一段文字", "http://")
	log.Printf("飞雪无情的公众号:%s\n", "fly_snow")
}
