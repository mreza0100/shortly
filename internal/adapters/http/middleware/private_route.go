package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
)

const UsernameKey = "username"

func PrivateRoute(jwt jwt.JWTHelper) gin.HandlerFunc {
	type ResponseBody struct {
		Error string `json:"error"`
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, ResponseBody{Error: "No token found"})
			return
		}

		email, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, ResponseBody{Error: "Invalid token"})
			return
		}

		ctx.Set(UsernameKey, email)
		ctx.Next()
	}
}
