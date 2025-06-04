package main

import (
	"fmt"
	"os"

	"r11manish.com/server"
)

func main() {
	host := "127.0.0.1"
	port := 8080
	fmt.Printf("::::::::::Http server::::::::::::")
	server := server.NewHttpServer(host, port)
	defer server.Cleanup()

	if err := server.Start(); err != nil {
		fmt.Printf("💥 Server error: %v\n", err)
		os.Exit(1)
	}
}
