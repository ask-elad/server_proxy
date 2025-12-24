package proxy

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/ask-elad/server_proxy/internal/filter"
	"github.com/ask-elad/server_proxy/internal/observ"
	"github.com/ask-elad/server_proxy/internal/protocol"
)

func Handle(client net.Conn, f *filter.Filter) {
	id := time.Now().UnixNano()
	log := func(msg string) {
		fmt.Printf("[conn %d] %s\n", id, msg)
	}

	defer func() {
		log("closing client connection")
		client.Close()
	}()

	reader := bufio.NewReader(client)

	result, err := protocol.Detect(reader)
	if err != nil {
		log("protocol detection failed: " + err.Error())
		return
	}
	var host string

	switch result.Kind {
	case protocol.HTTP:

		fields := strings.Fields(result.FirstLine)
		if len(fields) >= 2 {
			u, err := url.Parse(fields[1])
			if err == nil {
				host = u.Hostname()
			}
		}

	case protocol.CONNECT:

		fields := strings.Fields(result.FirstLine)
		if len(fields) >= 2 {
			host = strings.Split(fields[1], ":")[0]
		}
	}

	if f != nil && host != "" && f.IsBlocked(host) {
		log("blocked request to " + host)

		client.Write([]byte(
			"HTTP/1.1 403 Forbidden\r\n" +
				"Content-Length: 0\r\n\r\n",
		))

		observ.LogRequest(observ.RequestLog{
			Client: client.RemoteAddr().String(),
			Method: strings.Fields(result.FirstLine)[0],
			Target: host,
			Path:   "",
			Action: "BLOCK",
			Status: 403,
			Bytes:  0,
		})

		return
	}

	switch result.Kind {
	case protocol.HTTP:
		log("dispatching to HTTP handler")
		protocol.HandleHTTP(client, result)

	case protocol.CONNECT:
		log("dispatching to CONNECT handler")
		protocol.HandleCONNECT(client, result)

	default:
		log("unknown protocol, closing connection")
		return
	}
}
