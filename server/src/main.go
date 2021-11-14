package main

import "server"

func main() {
	server := server.GetServer()
	server.BindAndListen("8999")
}
