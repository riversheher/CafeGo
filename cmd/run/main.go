package main

import (
	"sync"

	"github.com/rainbowriverrr/CafeGo/pkg/database"
	"github.com/rainbowriverrr/CafeGo/pkg/server"
)

func main() {

	// Initializes the database and server configurations
	db := new(database.Database)
	server := new(server.Server)

	// Initializes the wait group and registers them with the server and database
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)

	db.WaitGroup = waitGroup

	server.WaitGroup = waitGroup

	go db.Start()
	go server.Start()

	waitGroup.Wait()

}
