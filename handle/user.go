package handle

import (
	"net/http"

	"example.com/gin_forum/logger"
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
	log := logger.New(ctx)
	// body := new(request.UserRegistrationRequest)
	// body:= &request.UserRegistrationRequest{}
	var body request.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithField("user", utils.JsonMarshal(body)).Info("user registration request received successfully")
}

func userLogin(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithField("user", utils.JsonMarshal(body)).Info("user login request received successfully")
}
