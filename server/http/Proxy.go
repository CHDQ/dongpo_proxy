package http

import (
	"bytes"
	"dongpo_proxy/proxy"
	"fmt"
	"io"
	"log"
	"net"
)

type Conn struct {
	method, host, httpVersion string
	client                    net.Conn
	server                    net.Conn
}

func StartServer(listenerAddr string) {
	serverController := &ServerController{}
	serverController.InitParam(proxy.Http, listenerAddr, "XOR", serverController)
	serverController.StartListen()
}

func handleRequest(connect net.Conn) {
	if connect == nil {
		return
	}
	var buffer [1024]byte
	num, err := connect.Read(buffer[:])
	if err != nil {
		log.Println(err)
		return
	}
	var con *Conn
	var method, host, httpVersion string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer[:num], '\n')]), "%s%s%s", &method, &host, &httpVersion)
	server, errors := net.Dial("tcp", host)
	if errors != nil {
		log.Println(errors)
		return
	}
	con = &Conn{
		method:      method,
		host:        host,
		httpVersion: httpVersion,
		client:      connect,
		server:      server,
	}
	if method == "CONNECT" { //处理开启隧道请求
		_, err := connect.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
	if con != nil {
		go io.Copy(con.client, con.server)
		io.Copy(con.server, con.client)
	}
}
