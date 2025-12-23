package protocol

import (
	"net"
	"strings"

	"github.com/ask-elad/server_proxy/internal/forwarder"
	"github.com/ask-elad/server_proxy/internal/observ"
	_ "github.com/ask-elad/server_proxy/internal/observ"
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

	bytesCopied, _ := forwarder.Tunnel(client, targetConn)

	observ.LogRequest(observ.RequestLog{
		Client: client.RemoteAddr().String(),
		Method: "CONNECT",
		Target: targetAddr,
		Path:   "",
		Action: "ALLOW",
		Status: 200,
		Bytes:  bytesCopied,
	})

}
