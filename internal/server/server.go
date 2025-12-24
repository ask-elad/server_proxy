package server

import (
	"fmt"
	"net"
	"time"

	"github.com/ask-elad/server_proxy/internal/config"
	"github.com/ask-elad/server_proxy/internal/filter"
	"github.com/ask-elad/server_proxy/internal/proxy"
)

func Run(cfg config.Config) error {

	ln, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("[server] listening on", cfg.ListenAddr)

	var f *filter.Filter
	if cfg.BlockedFile != "" {
		f, err = filter.Load(cfg.BlockedFile)
		if err != nil {
			return err
		}
		fmt.Println("[server] loaded blocklist from", cfg.BlockedFile)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[server] accept error:", err)
			continue
		}

		_ = conn.SetDeadline(time.Now().Add(cfg.ConnTimeout))

		fmt.Println("[server] accepted client:", conn.RemoteAddr())

		go proxy.Handle(conn, f)
	}
}
