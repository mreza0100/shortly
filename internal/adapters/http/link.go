package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/adapters/http/middleware"
	"github.com/mreza0100/shortly/internal/adapters/http/presenters"
	er "github.com/mreza0100/shortly/internal/pkg/errors"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
)

const linkParamKey = "link"

func registerLinkRoutes(ginClient *gin.Engine, jwtUtils jwt.JWTHelper, linkService ports.LinkServicePort) {
	linkHandlers := &linkHandlers{linkService: linkService}

	ginClient.GET("/:"+linkParamKey, linkHandlers.redirectByLink())

	group := ginClient.Group("/link")
	group.POST("/", middleware.PrivateRoute(jwtUtils), linkHandlers.newLink())
}

type linkHandlers struct {
	linkService ports.LinkServicePort
}

func (u *linkHandlers) newLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody presenters.NewLinkRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, presenters.NewLinkResponse{Error: err.Error()})
			return
		}

		short, err := u.linkService.NewLink(c.Request.Context(), requestBody.Link, "")
		if err != nil {
			c.JSON(er.Status(err), presenters.NewLinkResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, presenters.NewLinkResponse{Short: short})
	}
}

func (u *linkHandlers) redirectByLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortLink := c.Param(linkParamKey)
		destination, err := u.linkService.GetDestinationByLink(c.Request.Context(), shortLink)
		if err != nil {
			c.JSON(er.Status(err), presenters.RedirectByLinkResponse{Error: err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, destination)
	}
}
