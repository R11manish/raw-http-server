package server

import (
	"fmt"
	"strconv"

	"r11manish.com/model"
	"r11manish.com/utlis"
)

func (s *HttpServer) handleHome(req *model.HTTPRequest) *model.HTTPResponse {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Raw HTTP Server</title>
    <style>
        body { font-family: 'Courier New', monospace; margin: 40px; background: #1a1a1a; color: #00ff00; }
        .container { max-width: 900px; margin: 0 auto; }
        h1 { color: #00ff41; text-shadow: 0 0 10px #00ff41; }
        .endpoint { background: #2a2a2a; padding: 15px; margin: 10px 0; border-radius: 5px; border-left: 4px solid #00ff41; }
        .syscall { color: #ff6b6b; font-weight: bold; }
        .warning { background: #3a1a1a; color: #ffaa00; padding: 10px; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔥Raw HTTP Server</h1>
        <div class="warning">
            ⚡ This server uses raw syscalls - no net.Listen, no bufio, just pure system calls!
        </div>
        
        <h2>System Call Based Implementation:</h2>
        <p>This server uses these <span class="syscall">syscalls</span> directly:</p>
        <ul>
            <li><span class="syscall">syscall.Socket()</span> - Create TCP socket</li>
            <li><span class="syscall">syscall.Bind()</span> - Bind to address</li>
            <li><span class="syscall">syscall.Listen()</span> - Listen for connections</li>
            <li><span class="syscall">syscall.Accept()</span> - Accept connections</li>
            <li><span class="syscall">syscall.Read()</span> - Read data</li>
            <li><span class="syscall">syscall.Write()</span> - Write data</li>
            <li><span class="syscall">syscall.Close()</span> - Close sockets</li>
        </ul>
        
        <h2>Available Endpoints:</h2>
        <div class="endpoint">
            <strong>GET /</strong> - This home page
        </div>
        <div class="endpoint">
            <strong>POST /echo</strong> - Echo request details
        </div>
        
        <h2>Current Request:</h2>
        <p><strong>Method:</strong> ` + req.Method + `</p>
        <p><strong>Path:</strong> ` + req.Path + `</p>
        <p><strong>Version:</strong> ` + req.Version + `</p>
        <p><strong>User-Agent:</strong> ` + req.Headers["user-agent"] + `</p>
    </div>
</body>
</html>`

	return &model.HTTPResponse{
		Version:    "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type":   "text/html; charset=utf-8",
			"Content-Length": strconv.Itoa(len(html)),
			"Connection":     "close",
			"Server":         "UltraRaw/1.0",
		},
		Body: html,
	}
}

func (s *HttpServer) handleEcho(req *model.HTTPRequest) *model.HTTPResponse {
	response := fmt.Sprintf(`{
  "method": "%s",
  "path": "%s",
  "version": "%s",
  "headers": %s,
  "body": %s,
  "raw_syscall_server": true,
  "socket_fd": %d
}`, req.Method, req.Path, req.Version, utlis.FormatHeaders(req.Headers), strconv.Quote(req.Body), s.fd)

	return &model.HTTPResponse{
		Version:    "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"Content-Length": strconv.Itoa(len(response)),
			"Connection":     "close",
			"Server":         "UltraRaw/1.0",
		},
		Body: response,
	}
}

func (s *HttpServer) handle404(req *model.HTTPRequest) *model.HTTPResponse {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>404 Not Found - Ultra Raw Server</title>
    <style>
        body { font-family: 'Courier New', monospace; margin: 40px; text-align: center; background: #1a1a1a; color: #ff6b6b; }
        h1 { color: #ff4444; text-shadow: 0 0 10px #ff4444; }
        .syscall { color: #00ff41; }
    </style>
</head>
<body>
    <h1>404 Not Found</h1>
    <p>The path <strong>` + req.Path + `</strong> was not found.</p>
    <p>This error was delivered via <span class="syscall">raw syscalls</span>!</p>
    <p><a href="/" style="color: #00ff41;">Go back home</a></p>
</body>
</html>`

	return &model.HTTPResponse{
		Version:    "HTTP/1.1",
		StatusCode: 404,
		StatusText: "Not Found",
		Headers: map[string]string{
			"Content-Type":   "text/html; charset=utf-8",
			"Content-Length": strconv.Itoa(len(html)),
			"Connection":     "close",
			"Server":         "UltraRaw/1.0",
		},
		Body: html,
	}
}
