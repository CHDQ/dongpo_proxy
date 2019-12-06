package biz

import (
	"dongpo_proxy/proxy"
	"log"
	"net"
)

/**
 * 开启客户端监听
 */
func StartClient(listenerAddr string, rpcServerAddr string) {
	listener, err := net.ResolveTCPAddr("tcp", listenerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	rpcServer, err := net.ResolveTCPAddr("tcp", rpcServerAddr)
	xor := proxy.Xor{}
	conn := &proxy.Controller{Method: proxy.Http, Listener: listener, RpcConnector: rpcServer, XEncryption: xor}
	clientHandler := &ClientHandler{controller: *conn}
	conn.StartListen(clientHandler)
}
