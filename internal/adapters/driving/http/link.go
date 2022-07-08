package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/adapters/driving/http/middleware"
	"github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/pkg/jwt"
)

const linkParamKey = "link"

func registerLinkRoutes(ginClient *gin.Engine, jwtUtils jwt.JWTHelper, linkService services.LinkServicePort) {
	linkHandlers := &linkHandlers{linkService: linkService}

	ginClient.GET("/:"+linkParamKey, linkHandlers.redirectByLink())

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
			c.JSON(http.StatusBadRequest, ResponseBody{Error: err.Error()})
			return
		}

		short, err := u.linkService.NewLink(c.Request.Context(), requestBody.Link, "")
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseBody{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, ResponseBody{Short: short})
	}
}

func (u *linkHandlers) redirectByLink() gin.HandlerFunc {
	type ResponseBody struct {
		Error string `json:"error"`
	}

	return func(c *gin.Context) {
		shortLink := c.Param(linkParamKey)
		destination, err := u.linkService.GetDestinationByLink(c.Request.Context(), shortLink)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseBody{Error: err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, destination)
	}
}
