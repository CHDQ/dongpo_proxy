package main

import (
	"dongpo_proxy/client/biz"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start proxy client at " + time.Now().Format("2006-01-02 15:04:05"))
	biz.StartClient("127.0.0.1", "8100")
}