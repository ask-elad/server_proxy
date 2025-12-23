package server

import (
	"fmt"
	"net"
	"time"

	"github.com/ask-elad/server_proxy/internal/config"
	"github.com/ask-elad/server_proxy/internal/proxy"
)

// Run starts the proxy server using the provided configuration.
func Run(cfg config.Config) error {
	ln, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("[server] listening on", cfg.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[server] accept error:", err)
			continue
		}

		// Apply client connection lifetime timeout
		_ = conn.SetDeadline(time.Now().Add(cfg.ConnTimeout))

		fmt.Println("[server] accepted client:", conn.RemoteAddr())

		go proxy.Handle(conn)
	}
}
