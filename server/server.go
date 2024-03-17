package server

import (
	"example.com/gin_forum/handler"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	handler.AddUserHandler(r)
	handler.AddArticleHandler(r)
	handler.AddTagsHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
