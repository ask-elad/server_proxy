package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ask-elad/server_proxy/internal/proxy"
)

func main() {
	listenAddr := ":8080"
	targetAddr := "example.com:80"

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("[main] listening on", listenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[main] accept error:", err)
			continue
		}

		fmt.Println("[main] accepted client:", conn.RemoteAddr())
		go proxy.Handle(conn, targetAddr)
	}
}
