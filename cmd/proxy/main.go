package main

import (
	"fmt"
	"os"

	"github.com/ask-elad/server_proxy/internal/config"
	"github.com/ask-elad/server_proxy/internal/server"
)

func main() {
	cfg, err := config.Execute()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if err := server.Run(cfg); err != nil {
		fmt.Println("server error:", err)
		os.Exit(1)
	}
}
