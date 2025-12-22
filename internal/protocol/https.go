package protocol

import "net"

func HandleCONNECT(client net.Conn, res *Result) {
	// TODO:
	// - parse host:port from CONNECT line
	// - dial target
	// - send "200 Connection Established"
	// - switch to raw TCP forwarding
}
