package handle

import (
	"net/http"

	"example.com/gin_forum/logger"
	"example.com/gin_forum/models"
	"example.com/gin_forum/params/request"
	"example.com/gin_forum/params/response"
	"example.com/gin_forum/security"
	"example.com/gin_forum/storage"
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

	defaultUserImage := "https://api.realworld.io/images/smiley-cyrus.jpeg"

	if err := storage.CreateUser(ctx, &models.User{
		Username: body.User.Username,
		Password: body.User.Password,
		Email:    body.User.Email,
		Image:    defaultUserImage,
		Bio:      "",
	}); err != nil {
		log.WithError(err).Errorf("create user failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	token, err := security.GenerateJWT(body.User.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: body.User.Username,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		}})
	return
}

func userLogin(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithField("user", utils.JsonMarshal(body)).Info("user login request received successfully")
	dbUser, err := storage.GetUserByEmail(ctx, body.User.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO:
	if dbUser.Password != body.User.Password {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := security.GenerateJWT(dbUser.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: dbUser.Username,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		}})
}
