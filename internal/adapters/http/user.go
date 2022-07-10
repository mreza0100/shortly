package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/adapters/http/presenters"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
	"github.com/mreza0100/shortly/internal/ports"
)

func registerUserRoutes(ginClient *gin.Engine, userService ports.UserServicePort) {
	userHandlers := &userHandlers{userService: userService}
	group := ginClient.Group("/user")

	group.POST("/signup", userHandlers.signup())
	group.POST("/signin", userHandlers.signin())
}

type userHandlers struct {
	userService ports.UserServicePort
}

func (u *userHandlers) signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody presenters.SigninRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, presenters.SigninResponse{Error: err.Error()})
			return
		}

		if err := u.userService.Signup(c.Request.Context(), requestBody.Email, requestBody.Password); err != nil {
			c.JSON(customerror.Status(err), presenters.SigninResponse{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		c.Done()
	}
}

func (u *userHandlers) signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody presenters.SigninRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, presenters.SigninResponse{Error: err.Error()})
			return
		}

		token, err := u.userService.Signin(c.Request.Context(), requestBody.Email, requestBody.Password)
		if err != nil {
			c.JSON(customerror.Status(err), presenters.SigninResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, presenters.SigninResponse{Token: token})
	}
}
