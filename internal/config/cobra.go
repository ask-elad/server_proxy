package config

import (
	"errors"
	"time"

	"github.com/spf13/cobra"
)

func Execute() (Config, error) {
	cfg := Default()

	var (
		listenAddr  string
		dialTimeout time.Duration
		connTimeout time.Duration
		verbose     bool
	)

	rootCmd := &cobra.Command{
		Use:   "proxy",
		Short: "A minimal HTTP/HTTPS forward proxy",
		Long:  "A minimal, streaming, HTTP and HTTPS forward proxy written in Go.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if listenAddr != "" {
				cfg.ListenAddr = listenAddr
			}
			if dialTimeout > 0 {
				cfg.DialTimeout = dialTimeout
			}
			if connTimeout > 0 {
				cfg.ConnTimeout = connTimeout
			}
			cfg.Verbose = verbose

			if cfg.ListenAddr == "" {
				return errors.New("listen address cannot be empty")
			}

			return nil
		},
	}

	// Flags
	rootCmd.Flags().StringVar(
		&listenAddr,
		"listen",
		"",
		"Address to listen on (default :8080)",
	)

	rootCmd.Flags().DurationVar(
		&dialTimeout,
		"dial-timeout",
		0,
		"Timeout for outbound connections (e.g. 10s)",
	)

	rootCmd.Flags().DurationVar(
		&connTimeout,
		"conn-timeout",
		0,
		"Maximum lifetime of a client connection (e.g. 5m)",
	)

	rootCmd.Flags().BoolVar(
		&verbose,
		"verbose",
		false,
		"Enable verbose logging",
	)
	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
