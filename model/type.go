package model

import (
	"fmt"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

type HTTPResponse struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

func (r *HTTPResponse) ToBytes() []byte {
	var response strings.Builder

	response.WriteString(fmt.Sprintf("%s %d %s\r\n", r.Version, r.StatusCode, r.StatusText))

	for key, value := range r.Headers {
		response.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	response.WriteString("\r\n")
	response.WriteString(r.Body)
	return []byte(response.String())
}
