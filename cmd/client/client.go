package main

import (
	"keesvv/go-tcp-chat/internals/client"
	"keesvv/go-tcp-chat/internals/client/tui"
)

func main() {
	conn, err := client.Connect(client.Options{
		Hostname: "localhost",
		Port:     3300,
		Nickname: "keesvv",
	})

	if err != nil {
		panic(err)
	}

	go conn.HandleEvents()
	tui.Display(conn)
}
