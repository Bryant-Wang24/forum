package main

import (
	"example.com/gin_forum/cache"
	"example.com/gin_forum/server"
	_ "example.com/gin_forum/storage"
)

func main() {
	cache.InitRedis()
	server.RunHTTPServer()
}
