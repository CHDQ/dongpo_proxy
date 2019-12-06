package proxy

import (
	"io"
	"log"
	"net"
)

type Method string

const (
	Http   Method = "Http"
	SOCKS5 Method = "SOCKS5"
)

/**
代理管理
*/
type Controller struct {
	Method   Method       //http or socks5
	Listener *net.TCPAddr //本地监听

	XEncryption Encryption
	Handler
}

/**
加密
*/
type Encryption interface {
	Encode(input []byte) (int, []byte) //加密
	Decode(input []byte) (int, []byte) //解密
}

/**
处理网络请求
*/
type Handler interface {
	Handle(conn net.Conn) //进入时处理
}

/**
开启监听
*/
func (proxy *Controller) StartListen() {
	listener, err := net.ListenTCP("tcp", proxy.Listener)
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()
	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		localConn.SetLinger(0)
		go proxy.Handle(localConn)
	}
}
func (proxy *Controller) DialRpcServer(rpcServer *net.TCPAddr) (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, rpcServer)
	if err != nil {
		return nil, err
	}
	return remoteConn, nil
}

func (proxy *Controller) EncodeCopy(in net.Conn, out net.Conn) error {
	return proxy.copy(in, out, proxy.XEncryption.Encode)
}

func (proxy *Controller) DecodeCopy(in net.Conn, out net.Conn) error {
	return proxy.copy(in, out, proxy.XEncryption.Decode)
}

func (proxy *Controller) copy(in net.Conn, out net.Conn, handle func(data []byte) (int, []byte)) error {
	var buffer [1024]byte
	for {
		readCount, err := in.Read(buffer[:])
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
		readCount, data := handle(buffer[:])
		if readCount > 0 {
			writeCount, err := out.Write(data[:])
			if err != nil {
				return err
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}
func (proxy *Controller) InitParam(method Method, listenerAddr string, XEncryption string) {
	listener, err := net.ResolveTCPAddr("tcp", listenerAddr)
	if err != nil {
		log.Fatalln(err)
	}

	proxy.Method = method
	proxy.Listener = listener
	switch XEncryption {
	case "Xor":
		proxy.XEncryption = &Xor{}
		break
	default:
		proxy.XEncryption = &Xor{}
	}
}
