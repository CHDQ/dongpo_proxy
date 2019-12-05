package biz

import (
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
	server, err := net.Dial("tcp", "192.168.192.132:8080")
	if err != nil {
 		log.Fatal(err)
	}
	server.Write(buffer[:])
	go func() {
		fmt.Println(string(buffer[:]))
		server.Read(buffer[:])
		connect.Write(buffer[:])
		fmt.Println(string(buffer[:]))
	}()
	defer server.Close()
	defer connect.Close()
}
