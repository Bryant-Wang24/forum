package middlewares

import (
	"net/http"
	"strings"

	"example.com/gin_forum/logger"
	"example.com/gin_forum/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	log := logger.New(ctx)
	token := ctx.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims, ok, err := security.VerifyJWT(token)
	if err != nil || !ok {
		log.WithError(err).Infof("verify jwt failed")
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Set("user", claims)

	ctx.Next()
}
