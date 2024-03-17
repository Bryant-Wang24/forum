package handler

import (
	"net/http"
	"strings"

	"example.com/gin_forum/logger"
	"example.com/gin_forum/middlewares"
	"example.com/gin_forum/models"
	"example.com/gin_forum/params/request"
	"example.com/gin_forum/params/response"
	"example.com/gin_forum/security"
	"example.com/gin_forum/storage"
	"example.com/gin_forum/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddArticleHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	articlesGroup.GET("", listArticles)
	articlesGroup.GET("/:slug", getArticle)

	articlesGroup.Use(middlewares.AuthMiddleware)
	articlesGroup.POST("", createArticles)
	articlesGroup.PUT("/:slug", editArticles)
	articlesGroup.DELETE("/:slug", deleteArticles)
}

func createArticles(ctx *gin.Context) {
	//log := logger.New(ctx)
	var req request.CreateArticleRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.CreateArticle(ctx, &models.Article{
		AuthorUsername: security.GetCurrentUsername(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"article": respArticle,
	})
}

func editArticles(ctx *gin.Context) {
	//log := logger.New(ctx)
	oldSlug := ctx.Param("slug")
	var req request.CreateArticleRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	oldArticle, err := storage.GetArticleBySlug(ctx, oldSlug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUsername(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.UpdateArticle(ctx, oldSlug, &models.Article{
		AuthorUsername: security.GetCurrentUsername(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": respArticle,
	})
}

func deleteArticles(ctx *gin.Context) {
	slug := ctx.Param("slug")

	oldArticle, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUsername(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	err = storage.DeleteArticle(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func getArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": respArticle,
	})
}

func listArticles(ctx *gin.Context) {
	log := logger.New(ctx)
	var req request.ListArticleQuery
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Infof("list articles, req: %v\n", utils.JsonMarshal(req))

	articles, err := storage.ListArticles(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	total, err := storage.CountArticles(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var resp response.ListArticlesResponse
	resp.ArticlesCount = total
	for _, article := range articles {
		resp.Articles = append(resp.Articles, &response.Article{
			Author: &response.ArticleAuthor{
				Bio:       article.AuthorUserBio,
				Following: false,
				Image:     article.AuthorUserImage,
				Username:  article.AuthorUsername,
			},
			Title:          article.Title,
			Slug:           article.Slug,
			Body:           article.Body,
			Description:    article.Description,
			TagList:        article.TagList,
			Favorited:      false,
			FavoritesCount: 0,
			CreatedAt:      article.CreatedAt,
			UpdatedAt:      article.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, resp)

}
