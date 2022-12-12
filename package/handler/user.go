package handler

import (
	"api_getaway_web/client"
	"api_getaway_web/genproto/user_service"
	"api_getaway_web/model"
	"api_getaway_web/package/response"
	"api_getaway_web/tools/jwt"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.Engine, handler *Handler) {
	// Create routes group.
	user := api.Group("/api/v1/account")
	{
		user.POST("/sign-in", handler.SignInUser)
	}
}

// SignIn
// @Description User Sign In  user.
// @Summary User Sign In user
// @Tags User
// @Accept json
// @Produce json
// @Param signup body model.UserSignIn true "Sign In"
// @Success 200 {object} response.Response
// @Failure 400 {object}  response.Response
// @Failure 404 {object}  response.Response
// @Failure 500 {object}  response.Response
// @Router /api/v1/account/sign-in [post]
func (handler *Handler) SignInUser(ctx *gin.Context) {
	logrus := handler.logrus
	var body model.UserSignIn

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		logrus.Error(err)
		handler.handleResponse(ctx, response.BadRequest, err.Error())
		return
	}

	err = body.Validate()
	if err != nil {
		logrus.Error(err)
		handler.handleResponse(ctx, response.BadRequest, err.Error())
		return
	}

	body.Username = strings.TrimSpace(body.Username)
	body.Password = strings.TrimSpace(body.Password)
	user, serviceError := client.UserService().SigInUser(context.Background(), &user_service.SignInUserRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if serviceError != nil {
		logrus.Error(serviceError)
		handler.GrpcErrorConvert(ctx, serviceError)
		return
	}

	tokens, err := jwt.GenerateNewTokens(user.Id)
	if err != nil {
		logrus.Error(err)
		handler.handleResponse(ctx, response.InternalServerError, err.Error())
		return
	}

	handler.handleResponse(ctx, response.OK, tokens.Access)
}
