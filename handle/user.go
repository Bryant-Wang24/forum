package handle

import (
	"fmt"
	"net/http"

	"example.com/gin_forum/params/request"
	"example.com/gin_forum/utils"
	"github.com/gin-gonic/gin"
)

func AddUserHandler(r *gin.Engine) {
	userGroup := r.Group("/api/users")
	userGroup.POST("", userRegistration)
	userGroup.POST("/login", userLogin)
}

func userRegistration(ctx *gin.Context) {
	var body request.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(utils.JsonMarshal(body))
}

func userLogin(ctx *gin.Context) {
	var body request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(utils.JsonMarshal(body))
}
