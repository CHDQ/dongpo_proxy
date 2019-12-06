package http

import (
	"dongpo_proxy/proxy"
	"log"
	"net"
)

/**
 * 开启客户端监听
 */
func StartClient(listenerAddr string, rpcServerAddr string) {
	rpcServer, err := net.ResolveTCPAddr("tcp", rpcServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	clientController := &ClientController{RpcConnector: rpcServer}
	clientController.InitParam(proxy.Http, listenerAddr, "XOR", clientController)
	clientController.StartListen()
}
