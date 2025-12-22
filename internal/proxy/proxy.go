package proxy

import (
	"fmt"
	"io"
	"net"
	"time"
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

	log("connecting to target " + targetAddr)

	target, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log("failed to connect to target: " + err.Error())
		return
	}
	defer func() {
		log("closing target connection")
		target.Close()
	}()

	clientTCP, _ := client.(*net.TCPConn)
	targetTCP, _ := target.(*net.TCPConn)

	log("starting client -> target copy")
	go func() {
		n, err := io.Copy(target, client)
		log(fmt.Sprintf("client -> target finished (%d bytes, err=%v)", n, err))

		if targetTCP != nil {
			log("closing target write side")
			targetTCP.CloseWrite()
		}
	}()

	log("starting target -> client copy")
	n, err := io.Copy(client, target)
	log(fmt.Sprintf("target -> client finished (%d bytes, err=%v)", n, err))

	if clientTCP != nil {
		log("closing client write side")
		clientTCP.CloseWrite()
	}

	log("handler exiting")
}
