package proxy

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/ask-elad/server_proxy/internal/protocol"
)

func Handle(client net.Conn, targetAddr string) {
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
