package config

import "time"

type Config struct {
	ListenAddr  string
	DialTimeout time.Duration
	ConnTimeout time.Duration
	Verbose     bool
}

func Default() Config {
	return Config{
		ListenAddr:  ":8080",
		DialTimeout: 10 * time.Second,
		ConnTimeout: 5 * time.Minute,
		Verbose:     false,
	}
}
