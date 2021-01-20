package main

import (
	"fmt"
	"keesvv/go-tcp-chat/internals/server"
)

func main() {
	listener := server.Listener{
		Bind: []string{"127.0.0.1", "::1"},
		Port: 3300,
	}

	err := listener.Listen()
	if err != nil {
		panic(err)
	}

	// Exit on return
	fmt.Scanln()
}