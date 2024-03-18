package handler

import (
	"net/http"

	"example.com/gin_forum/cache"
	"example.com/gin_forum/storage"
	"github.com/gin-gonic/gin"
)

func AddTagsHandler(r *gin.Engine) {
	tagsGroup := r.Group("/api/tags")
	tagsGroup.GET("", listPopularTags)
}

func listPopularTags(ctx *gin.Context) {
	tags, _ := cache.GetPopularTags(ctx)
	if len(tags) != 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"tags": tags,
		})
		return
	}
	tags, err := storage.ListPopularTags(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"tags": tags,
	})
}
