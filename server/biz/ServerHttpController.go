package biz

import (
	"bytes"
	"dongpo_proxy/proxy"
	"fmt"
	"log"
	"net"
)

type ServerController struct {
	proxy.Controller
}

func (serverController *ServerController) Handle(conn net.Conn) {
	defer conn.Close()
	var buffer [1024]byte
	num, err := conn.Read(buffer[:])
	if err != nil {
		log.Fatal(err)
	}
	var method, host, httpVersion string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer[:num], '\n')]), "%s%s%s", &method, &host, &httpVersion)
	rpcServer, errors := net.Dial("tcp", host)
	defer rpcServer.Close()
	if errors != nil {
		log.Fatal(errors)
		return
	}
	if method == "CONNECT" { //处理开启隧道请求
		_, data := serverController.XEncryption.Encode([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		_, err :=rpcServer.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
	}
	go func() {
		serverController.DecodeCopy(rpcServer, conn)
	}()
	serverController.EncodeCopy(conn, rpcServer)
}
