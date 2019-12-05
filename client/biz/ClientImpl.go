package biz

import (
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
	defer connect.Close()
	var buffer [1024]byte
	_, err := connect.Read(buffer[:])
	if err != nil {
		log.Fatal(err)
	}

}
