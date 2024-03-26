package v1

import (
	"net/http"
	"xsis-movie/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, t usecase.Movies) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	h := handler.Group("/movies")
	{
		newMovieRoutes(h, t)
	}
}
