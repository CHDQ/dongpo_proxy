package main

import (
	"dongpo_proxy/server/http"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start proxy server at " + time.Now().Format("2006-01-02 15:04:05"))
	http.StartServer(":8080")
}
