package protocol

import (
	"io"
	"net"
	"strings"
)

func HandleCONNECT(client net.Conn, res *Result) {

	fields := strings.Fields(res.FirstLine)
	if len(fields) != 3 {
		return
	}

	targetAddr := fields[1]
	if !strings.Contains(targetAddr, ":") {
		return
	}

	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		return
	}

	clientTCP, _ := client.(*net.TCPConn)
	targetTCP, _ := targetConn.(*net.TCPConn)

	_, err = client.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	if err != nil {
		targetConn.Close()
		return
	}

	go func() {
		_, _ = io.Copy(targetConn, res.Reader)
		if targetTCP != nil {
			targetTCP.CloseWrite()
		}
	}()

	_, _ = io.Copy(client, targetConn)
	if clientTCP != nil {
		clientTCP.CloseWrite()
	}
}
