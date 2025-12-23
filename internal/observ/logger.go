package observ

import (
	"log"
	"time"
)

type RequestLog struct {
	Client string
	Method string
	Target string
	Path   string
	Action string
	Status int
	Bytes  int64
}

func LogRequest(r RequestLog) {
	ts := time.Now().UTC().Format(time.RFC3339)

	log.Printf(
		"[%s] client=%s method=%s target=%s path=%s action=%s status=%d bytes=%d",
		ts,
		r.Client,
		r.Method,
		r.Target,
		r.Path,
		r.Action,
		r.Status,
		r.Bytes,
	)
}
