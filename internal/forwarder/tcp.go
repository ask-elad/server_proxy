package forwarder

import (
	"io"
	"net"
)

func Tunnel(a, b net.Conn) {
	aTCP, _ := a.(*net.TCPConn)
	bTCP, _ := b.(*net.TCPConn)

	go func() {
		_, _ = io.Copy(b, a)
		if bTCP != nil {
			bTCP.CloseWrite()
		}
	}()

	_, _ = io.Copy(a, b)
	if aTCP != nil {
		aTCP.CloseWrite()
	}
}
