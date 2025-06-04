package server

import (
	"fmt"
	"strconv"
	"syscall"

	"r11manish.com/model"
	"r11manish.com/utlis"
)

func (s *HttpServer) acceptConnection() (int, *syscall.SockaddrInet4, error) {
	clientFd, clientAddr, err := syscall.Accept(s.fd)

	if err != nil {
		return 0, nil, fmt.Errorf("failed to accept connection : %v", err)
	}

	addr, ok := clientAddr.(*syscall.SockaddrInet4)

	if !ok {
		syscall.Close(clientFd)
		return 0, nil, fmt.Errorf("unexpected address type")
	}

	return clientFd, addr, nil
}

func (s *HttpServer) readFromSocket(fd int) ([]byte, error) {
	buffer := make([]byte, 4096)
	n, err := syscall.Read(fd, buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read from socket: %v", err)
	}

	return buffer[:n], nil
}

func (s *HttpServer) writeToSocket(fd int, data []byte) error {
	totalWritten := 0
	for totalWritten < len(data) {
		n, err := syscall.Write(fd, data[totalWritten:])

		if err != nil {
			return fmt.Errorf("failed to write to socket : %v", err)
		}
		totalWritten += n
	}
	return nil
}

func (s *HttpServer) handleConnection(clientFd int, clientAddr *syscall.SockaddrInet4) {
	defer s.closeSocket(clientFd)

	clientIp := fmt.Sprintf("%d.%d.%d.%d", clientAddr.Addr[0], clientAddr.Addr[1], clientAddr.Addr[2], clientAddr.Addr[3])

	fmt.Printf("New connection from %s:%d (fd: %d)\n", clientIp, clientAddr.Port, clientFd)

	timeout := syscall.Timeval{Sec: 30, Usec: 0}

	syscall.SetsockoptTimeval(clientFd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &timeout)
	syscall.SetsockoptTimeval(clientFd, syscall.SOL_SOCKET, syscall.SO_SNDTIMEO, &timeout)

	requestData, err := s.readFromSocket(clientFd)
	if err != nil {
		fmt.Printf("Error reading request : %v \n", err)
		return
	}

	request, err := utlis.ParsedHttpRequest(requestData)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		errorResponse := s.buildErrorResponse(400, "Bad Request")
		s.writeToSocket(clientFd, errorResponse.ToBytes())
		return
	}

	fmt.Printf("Message: %s %s from %s\n", request.Method, request.Path, clientIp)

	response := s.routeRequest(request)
	responseBytes := response.ToBytes()
	err = s.writeToSocket(clientFd, responseBytes)
	if err != nil {
		fmt.Printf("Error sending response: %v\n", err)
	} else {
		fmt.Printf("Response sent (%d bytes)\n", len(responseBytes))
	}
}

func (s *HttpServer) routeRequest(req *model.HTTPRequest) *model.HTTPResponse {
	switch {
	case req.Method == "GET" && req.Path == "/":
		return s.handleHome(req)
	case req.Method == "POST" && req.Path == "/echo":
		return s.handleEcho(req)
	default:
		return s.handle404(req)
	}
}

func (s *HttpServer) buildErrorResponse(statusCode int, statusText string) *model.HTTPResponse {
	body := fmt.Sprintf(`{"error": %d, "message": "%s", "server": "ultra_raw_syscalls"}`, statusCode, statusText)
	return &model.HTTPResponse{
		Version:    "HTTP/1.1",
		StatusCode: statusCode,
		StatusText: statusText,
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"Content-Length": strconv.Itoa(len(body)),
			"Connection":     "close",
			"Server":         "RawServer/1.0",
		},
		Body: body,
	}
}
