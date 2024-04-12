package main

import (
	"example.com/gin_forum/cache"
	"example.com/gin_forum/server"
	_ "example.com/gin_forum/storage"
)

// 添加热加载工具fresh
//参考 https://github.com/gravityblast/fresh/issues/74#issuecomment-1762715817

func main() {
	cache.InitRedis()
	server.RunHTTPServer()
}
