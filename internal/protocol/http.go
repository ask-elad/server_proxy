package protocol

import "net"

func HandleHTTP(client net.Conn, res *Result) {
	// TODO:
	// - parse absolute-form request line
	// - extract host:port
	// - rewrite request line
	// - forward request + response
}
