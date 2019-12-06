package biz

import (
	"dongpo_proxy/proxy"
	"log"
	"net"
)

type ClientController struct {
	proxy.Controller
	RpcConnector *net.TCPAddr //远程调用
}

func (clientController *ClientController) Handle(conn net.Conn) {
	defer conn.Close()
	rpcServer, err := clientController.DialRpcServer(clientController.RpcConnector)
	defer rpcServer.Close()
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		clientController.DecodeCopy(rpcServer, conn)
	}()
	clientController.EncodeCopy(conn, rpcServer)
}
