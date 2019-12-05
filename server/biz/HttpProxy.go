package biz

import (
	"bytes"
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

func StartServer(ip string, port string) {
	listen, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatal("Server start error!!!\n", err)
		return
	}
	log.Print("Server has started. Listening "+"address : ", ip+":"+port+"	...")
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
	num, err := connect.Read(buffer[:])
	if err != nil {
		log.Fatal(err)
	}
	var con *Conn
	var method, host, httpVersion string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer[:num], '\n')]), "%s%s%s", &method, &host, &httpVersion)
	server, errors := net.Dial("tcp", host)
	if errors != nil {
		log.Fatal(errors)
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
			log.Fatal(err)
		}
	}
	if con != nil {
		go io.Copy(con.client, con.server)
		io.Copy(con.server, con.client)
	}
}
