package protocol

import (
	"io"
	"net"
	"net/url"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Version string
	Target  string
	Path    string
}

func HandleHTTP(client net.Conn, res *Result) {

	fields := strings.Fields(res.FirstLine)
	if len(fields) != 3 {
		return
	}

	method := fields[0]
	rawURL := fields[1]
	version := fields[2]

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return
	}

	if parsedURL.Scheme != "http" {
		return
	}

	host := parsedURL.Hostname()
	if host == "" {
		return
	}

	port := parsedURL.Port()
	if port == "" {
		port = "80"
	}

	path := parsedURL.RequestURI()
	if path == "" {
		path = "/"
	}

	targetAddr := host + ":" + port

	req := HTTPRequest{
		Method:  method,
		Version: version,
		Target:  targetAddr,
		Path:    path,
	}

	targetConn, err := net.Dial("tcp", req.Target)
	if err != nil {
		return
	}

	clientTCP, _ := client.(*net.TCPConn)
	targetTCP, _ := targetConn.(*net.TCPConn)

	rewrittenFirstLine := req.Method + " " + req.Path + " " + req.Version + "\r\n"
	_, err = targetConn.Write([]byte(rewrittenFirstLine))
	if err != nil {
		targetConn.Close()
		return
	}

	go func() {
		_, _ = io.Copy(targetConn, res.Reader)
		if targetTCP != nil {
			targetTCP.CloseWrite()
		}
	}()

	_, _ = io.Copy(client, targetConn)

	if clientTCP != nil {
		clientTCP.CloseWrite()
	}
}
