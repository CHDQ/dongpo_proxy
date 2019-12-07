package http

import (
	"dongpo_proxy/proxy"
	"net"
)

var clientController *ClientController

/**
 * 开启客户端监听
 */
func StartClient(listenerAddr string, rpcServerAddr string) error {
	rpcServer, err := net.ResolveTCPAddr("tcp", rpcServerAddr)
	if err != nil {
		return err
	}
	clientController = &ClientController{RpcConnector: rpcServer}
	clientController.InitParam(proxy.Http, listenerAddr, "XOR", clientController)
	return clientController.StartListen()
}

func ShutdownClient() {
	if clientController != nil {
		clientController.IsStop = true
		clientController.ShutdownListener()
	}
}
