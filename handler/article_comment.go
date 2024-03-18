package handler

import (
	"net/http"

	"example.com/gin_forum/logger"
	"example.com/gin_forum/middlewares"
	"example.com/gin_forum/models"
	"example.com/gin_forum/params/response"
	"example.com/gin_forum/security"
	"example.com/gin_forum/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func AddArticleCommentHandler(r *gin.Engine) {
	commentsGroup := r.Group("/api/articles/:slug/comments")
	commentsGroup.GET("", getArticleComments)

	commentsGroup.Use(middlewares.AuthMiddleware)
	commentsGroup.POST("", createArticleComment)
	commentsGroup.DELETE("/:commentId", deleteArticleComment)
}

func getArticleComments(ctx *gin.Context) {
	log := logger.New(ctx)
	slug := ctx.Param("slug")
	log.Infof("get article comments, slug: %v", slug)

	articleComments, err := storage.GetArticleCommentsByArticleId(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var resp []*response.ArticleComment
	for _, articleComment := range articleComments {
		respArticleComment := &response.ArticleComment{}
		respArticleComment.FromDB(articleComment)
		resp = append(resp, respArticleComment)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"comments": resp,
	})
}

func createArticleComment(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//req := make(map[string]interface{})
	req := struct {
		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`
	}{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	creator := security.GetCurrentUsername(ctx)

	articleComment := &models.ArticleComment{
		AuthorUsername: creator,
		Body:           req.Comment.Body,
		ArticleId:      article.Id,
	}
	if err := storage.CreateArticleComment(ctx, articleComment); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	comment, err := storage.GetArticleCommentById(ctx, articleComment.Id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &response.ArticleComment{}
	resp.FromDB(comment)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"comment": resp,
	})
}

func deleteArticleComment(ctx *gin.Context) {
	commentId := cast.ToInt64(ctx.Param("commentId"))

	articleComment, err := storage.GetArticleCommentById(ctx, commentId)
	if err != nil {
		if storage.IsNotFound(err) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	currentUsername := security.GetCurrentUsername(ctx)
	if currentUsername != articleComment.AuthorUsername {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err := storage.DeleteArticleCommentById(ctx, commentId); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
