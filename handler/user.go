package handler

import (
	"net/http"

	"example.com/gin_forum/cache"
	"example.com/gin_forum/logger"
	"example.com/gin_forum/middlewares"
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
	r.GET("/api/profiles/:username", userProfile)
	r.Group("/api/user").Use(middlewares.AuthMiddleware).PUT("", editUser)
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

	hashPassword, err := security.HashPassword(body.User.Password)
	if err != nil {
		log.WithError(err).Errorf("hash password failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := storage.CreateUser(ctx, &models.User{
		Username: body.User.Username,
		Password: hashPassword,
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

	if !security.CheckPassword(body.User.Password, dbUser.Password) {
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

func userProfile(ctx *gin.Context) {
	log := logger.New(ctx)
	userName := ctx.Param("username")
	log = log.WithField("username", userName)
	log.Infof("user profile called, userName: %v\n", userName)

	var user *models.User
	user, _ = cache.GetUserProfile(ctx, userName)
	if user == nil {
		var err error
		user, err = storage.GetUserByUsername(ctx, userName)
		if err != nil {
			log.WithError(err).Infoln("get user by username failed")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := cache.SetUserProfile(ctx, userName, user, 300); err != nil {
			log.WithError(err).Infoln("set user profile failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, response.UserProfileResponse{UserProfile: response.UserProfile{
		Username:  user.Username,
		Bio:       user.Bio,
		Image:     user.Image,
		Following: false,
	}})
}

func editUser(ctx *gin.Context) {
	log := logger.New(ctx)
	log.Infof("edit user: %v", security.GetCurrentUsername(ctx))
	var body request.EditUserRequest
	if err := ctx.BindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if body.EditUserBody.Username == "" || body.EditUserBody.Email == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if body.EditUserBody.Password != "" {
		var err error
		body.EditUserBody.Password, err = security.HashPassword(body.EditUserBody.Password)
		if err != nil {
			log.WithError(err).Errorln("hash password failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	dbUser := &models.User{
		Username: body.EditUserBody.Username,
		Password: body.EditUserBody.Password,
		Email:    body.EditUserBody.Email,
		Image:    body.EditUserBody.Image,
		Bio:      body.EditUserBody.Bio,
	}
	if err := storage.UpdateUserByUsername(ctx, security.GetCurrentUsername(ctx), dbUser); err != nil {
		log.WithError(err).Errorf("UpdateUserByUsername failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := cache.DeleteUserProfile(ctx, security.GetCurrentUsername(ctx)); err != nil {
		log.WithError(err).Error("DeleteUserProfile failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := security.GenerateJWT(dbUser.Username, dbUser.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    dbUser.Email,
			Token:    token,
			Username: dbUser.Username,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		}})
}
