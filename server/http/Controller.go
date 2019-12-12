package http

import (
	"bytes"
	"dongpo_proxy/proxy"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
)

type ServerController struct {
	proxy.Controller
}

func (serverController *ServerController) Handle(conn net.Conn) {
	defer conn.Close()
	var buffer [1024]byte
	num, err := conn.Read(buffer[:])
	if err != nil {
		log.Println(err)
		return
	}
	num, data := serverController.XEncryption.Decode(buffer[:])
	if num > 0 && bytes.IndexByte(data[:num], '\n') > 0 {
		log.Println(string(data[:num]))
		var method, host, httpVersion string
		fmt.Sscanf(string(data[:bytes.IndexByte(data[:num], '\n')]), "%s%s%s", &method, &host, &httpVersion)
		hostPortURL, err := url.Parse(host)
		var address string
		if err != nil {
			log.Println(err)
			return
		}
		if hostPortURL.Opaque == "443" { //https访问
			address = hostPortURL.Scheme + ":443"
		} else {                                            //http访问
			if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
				address = hostPortURL.Host + ":80"
			} else {
				address = hostPortURL.Host
			}
		}
		rpcServer, errors := net.Dial("tcp", address)
		if errors != nil {
			log.Println(errors)
			return
		}
		defer rpcServer.Close()
		if method == "CONNECT" { //处理开启隧道请求
			_, data := serverController.XEncryption.Encode([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
			_, err := conn.Write(data)
			if err != nil {
				log.Println(err)
				return
			}
		}
		go func() {
			err := serverController.DecodeCopy(rpcServer, conn)
			if err != nil {
				log.Println(err)
			}
		}()
		err = serverController.EncodeCopy(conn, rpcServer)
		if err != nil {
			log.Println(err)
		}
	}
}
