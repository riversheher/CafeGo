package main

import (
	"github.com/rainbowriverrr/CafeGo/pkg/server"
)

func main() {

	// Initializes the server configurations
	server := new(server.Server)

	server.Start()

}
