package main

import (
	"dongpo_proxy/client/http"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start proxy client at " + time.Now().Format("2006-01-02 15:04:05"))
	http.StartClient(":8100", "167.179.102.249:8080")
	//http.StartClient(":8100", "127.0.0.1:8080")
}
