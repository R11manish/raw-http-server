package server

import (
	"fmt"
	"syscall"
)

type HttpServer struct {
	host string
	port int
	fd   int
}

func NewHttpServer(host string, port int, fd int) *HttpServer {
	return &HttpServer{
		host: host,
		port: port,
	}
}

func (s *HttpServer) createSocket() error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)

	if err != nil {
		return fmt.Errorf("failed to create socket : %v", err)
	}

	s.fd = fd
	fmt.Printf("Created socket with file descriptor: %d\n", fd)
	return nil
}

func (s *HttpServer) setSocketOptions() error {
	err := syscall.SetsockoptInt(s.fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)

	if err != nil {
		return fmt.Errorf("failed to set SO_REUSEADDR: %v", err)
	}

	err = syscall.SetsockoptInt(s.fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)

	if err != nil {
		fmt.Printf("warning: failed to set SO_REUSEPORT: %v\n", err)
	}

	fmt.Println("Socket option set successfully")

	return nil
}

func (s *HttpServer) bindSocket() error {
	ip := [4]byte{127, 0, 0, 1}
	if s.host != "localhost" && s.host != "127.0.0.1" {
		fmt.Printf(" Using localhost instead of %s\n", s.host)
	}

	addr := syscall.SockaddrInet4{
		Port: s.port,
		Addr: ip,
	}

	err := syscall.Bind(s.fd, &addr)
	if err != nil {
		return fmt.Errorf("failed to bind socket to %s:%d: %v", s.host, s.port, err)
	}

	fmt.Printf(" Socket bound to %s:%d\n", s.host, s.port)
	return nil
}

func (s *HttpServer) listenSocket() error {
	err := syscall.Listen(s.fd, 128)
	if err != nil {
		return fmt.Errorf("failed to listen on socket: %v", err)
	}

	fmt.Println("socket listening for connections")
	return nil
}

func (s *HttpServer) closeSocket(fd int) {
	syscall.Close(fd)
}

func (s *HttpServer) Start() error {
	fmt.Printf(">>=====Starting http server========>>")

	if err := s.createSocket(); err != nil {
		return err
	}
	defer s.closeSocket(s.fd)

	if err := s.setSocketOptions(); err != nil {
		return err
	}

	if err := s.bindSocket(); err != nil {
		return err
	}

	if err := s.listenSocket(); err != nil {
		return err
	}

	fmt.Printf("server started listening")

	// write a loop
	// main server loop

}

func (s *HttpServer) Cleanup() {
	if s.fd > 0 {
		s.closeSocket(s.fd)
		fmt.Println("server socket closed")
	}
}

