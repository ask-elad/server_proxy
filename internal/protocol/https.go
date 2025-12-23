package protocol

import (
	"net"
	"strings"

	"github.com/ask-elad/server_proxy/internal/forwarder"
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

	_, err = client.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	if err != nil {
		targetConn.Close()
		return
	}

	forwarder.Tunnel(client, targetConn)
}
