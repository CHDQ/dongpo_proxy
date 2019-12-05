package biz

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func startServer() {

}

func DoRequest(buffer []byte, client net.Conn) {
	var method, host string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
		return
	}
	var address string
	if hostPortURL.Opaque == "443" { //https请求
		address = hostPortURL.Scheme + ":443"
	} else {                                            //http请求
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	fmt.Print(string(buffer[:]))
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Write(buffer[:])
	client.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

	// 直通双向复制
	go io.Copy(server, client)
	go io.Copy(client, server)
}
