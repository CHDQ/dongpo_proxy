package windows

import (
	"golang.org/x/sys/windows/registry"
	"strings"
)

const (
	ProxyOn     byte = 0x02
	PacOn       byte = 0x04
	AutoFoundOn byte = 0x08
)

/**
windows设置网络代理
*/
func SetProxy(proxyAddr string, pacAddr string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings\Connections`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	code := generateProxyBytes(proxyAddr, "", pacAddr, false)
	return key.SetBinaryValue(`DefaultConnectionSettings`, code)
}

/**
关闭代理
*/
func CloseProxy() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings\Connections`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	value, _, err := key.GetBinaryValue(`DefaultConnectionSettings`)
	if err != nil {
		return err
	}
	value[8] = 0x01
	return key.SetBinaryValue(`DefaultConnectionSettings`, value)
}

/**
生成修改注册表字节的字节码
*/
func generateProxyBytes(proxyAddr string, localhost string, pacAddr string, isAutoFound bool) []byte {
	var code = []byte{0x46, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	var nullByte = []byte{0x00, 0x00, 0x00}
	code = append(code, nullByte...)
	proxyAddr = strings.TrimSpace(proxyAddr)
	localhost = strings.TrimSpace(localhost)
	pacAddr = strings.TrimSpace(pacAddr)
	code = append(code, byte(len(proxyAddr)))
	code = append(code, nullByte...)
	if proxyAddr != "" {
		code[8] = code[8] | ProxyOn
		code = append(code, []byte(proxyAddr)...)
	}
	code = append(code, byte(len(localhost)))
	code = append(code, nullByte...)
	if localhost != "" {
		code = append(code, []byte(localhost)...)
		code = append(code, 0x3b, 0x3c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x3e) //local
	}
	code = append(code, byte(len(pacAddr)))
	code = append(code, nullByte...)
	if pacAddr != "" {
		code[8] = code[8] | PacOn
		code = append(code, []byte(pacAddr)...)
	}
	code = append(code, 0x01)
	if isAutoFound {
		code[8] = code[8] | AutoFoundOn
	}
	return code
}
