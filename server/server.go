package server

import (
	"time"

	"example.com/gin_forum/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	handler.AddUserHandler(r)
	handler.AddArticleHandler(r)
	handler.AddTagsHandler(r)
	handler.AddArticleCommentHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
