package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/adapters/driving/http/middleware"
	"github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/pkg/jwt"
)

func registerLinkRoutes(ginClient *gin.Engine, jwtUtils jwt.JWTHelper, linkService services.LinkServicePort) {
	linkHandlers := &linkHandlers{linkService: linkService}
	group := ginClient.Group("/link")

	group.POST("/", middleware.PrivateRoute(jwtUtils), linkHandlers.newLink())
}

type linkHandlers struct {
	linkService services.LinkServicePort
}

func (u *linkHandlers) newLink() gin.HandlerFunc {
	type RequestBody struct {
		Link string `json:"link"`
	}
	type ResponseBody struct {
		Short string `json:"short"`
		Error string `json:"error"`
	}
	return func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}

		short, err := u.linkService.NewLink(c.Request.Context(), requestBody.Link, "")
		if err != nil {
			c.JSON(400, ResponseBody{Error: err.Error()})
			return
		}

		c.JSON(200, ResponseBody{Short: short})
	}
}
