package protocol

import (
	"bufio"
	"strings"
)

type Kind int

const (
	Unknown Kind = iota
	HTTP
	CONNECT
)

type Result struct {
	Kind      Kind
	FirstLine string
	Reader    *bufio.Reader
}

func Detect(reader *bufio.Reader) (*Result, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)

	switch {
	case strings.HasPrefix(line, "CONNECT"):
		return &Result{
			Kind:      CONNECT,
			FirstLine: line,
			Reader:    reader,
		}, nil

	case strings.HasPrefix(line, "GET"),
		strings.HasPrefix(line, "POST"),
		strings.HasPrefix(line, "HEAD"),
		strings.HasPrefix(line, "PUT"),
		strings.HasPrefix(line, "DELETE"):
		return &Result{
			Kind:      HTTP,
			FirstLine: line,
			Reader:    reader,
		}, nil

	default:
		return &Result{
			Kind:      Unknown,
			FirstLine: line,
			Reader:    reader,
		}, nil
	}
}
