package biz

import (
	"dongpo_proxy/proxy"
	"io"
	"log"
	"net"
)

type ClientHandler struct {
	controller proxy.Controller
}

func (clientHandler *ClientHandler) LocalHandle(conn net.Conn) {
	defer conn.Close()
	var buffer [1024]byte
	rpcServer, err := clientHandler.controller.DialRpcServer()
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		rpcServer
	}()

}
func (clientHandler *ClientHandler) RpcHandle(conn net.Conn) {

}

