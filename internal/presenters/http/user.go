package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/ports/services"
)

func registerUserRoutes(ginClient *gin.Engine, userService services.UserServicePort) {
	userHandlers := &userHandlers{userService: userService}
	group := ginClient.Group("/user")

	group.POST("/signup", userHandlers.signup())
	group.POST("/signin", userHandlers.signin())
}

type userHandlers struct {
	userService services.UserServicePort
}

func (u *userHandlers) signup() gin.HandlerFunc {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type ResponseBody struct {
		Error string `json:"error"`
	}

	return func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}

		err := u.userService.Signup(requestBody.Email, requestBody.Password)
		if err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}
	}
}

func (u *userHandlers) signin() gin.HandlerFunc {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type ResponseBody struct {
		Token string `json:"token"`
		Error string `json:"error"`
	}

	return func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}

		token, err := u.userService.Signin(requestBody.Email, requestBody.Password)
		if err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}

		c.JSON(200, ResponseBody{Token: token})
	}
}
