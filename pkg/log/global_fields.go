package log

import (
	"strings"

	"github.com/rs/zerolog"
)

type GlobalLogFields struct {
	ID         string
	RemoteIP   string
	Host       string
	Method     string
	Path       string
	Protocol   string
	StatusCode int
	Latency    float64
	Error      error
	Stack      []byte
}

func (f *GlobalLogFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("id", f.ID).
		Str("remote_ip", f.RemoteIP).
		Str("host", f.Host).
		Str("method", strings.TrimSpace(f.Method)).
		Str("path", f.Path).
		Str("protocol", f.Protocol).
		Int("status_code", f.StatusCode).
		Float64("latency", f.Latency).
		Str("tag", "request")

	if f.Error != nil {
		e.Err(f.Error)
	}

	if f.Stack != nil {
		e.Bytes("stack", f.Stack)
	}
}
