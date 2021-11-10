package main

import "server"

func main() {
	server := server.GetServer()
	server.Listen("8999")
}
