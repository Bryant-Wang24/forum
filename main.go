package main

import (
	"example.com/gin_forum/server"
	_ "example.com/gin_forum/storage"
)

func main() {
	server.RunHTTPServer()
}
