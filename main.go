package main

import "github.com/neocpx/gochat/backend"

func main() {
	server := backend.NewHost("8080")
	server.Run()
}
