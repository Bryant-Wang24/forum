package server

import (
	"example.com/gin_forum/handle"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	handle.AddUserHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
