package biz

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

/**
 * 开启客户端监听
 */
func StartClient(ip string, port string) {
	listen, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatal("client start error!!!\n", err)
		return
	}
	log.Print("Client has started. Listening "+"address : ", ip+":"+port+"	...")
	for {
		connect, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("access connect [" + connect.RemoteAddr().String() + "]")
		go handleRequest(connect)
	}
}

func handleRequest(connect net.Conn) {
	if connect == nil {
		return
	}
	var buffer [1024]byte
	_, err := connect.Read(buffer[:])
	if err != nil {
		log.Fatal(err)
	}
	var method, host, httpVersion string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer[:], '\n')]), "%s%s%s", &method, &host, &httpVersion)
	if method == "CONNECT" { //处理请求
		num, err := connect.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(num)
	}
	for {
		num, _ := connect.Read(buffer[:])
		fmt.Println(num)
		fmt.Println(string(buffer[:]))
	}
	//biz.DoRequest(buffer[:num], connect)
}
