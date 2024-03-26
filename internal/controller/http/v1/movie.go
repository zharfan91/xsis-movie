package v1

import (
	"log"
	"net/http"
	"xsis-movie/internal/entity"
	"xsis-movie/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MovieRouters struct {
	m usecase.Movies
}

func newMovieRoutes(handler *gin.RouterGroup, t usecase.Movies) {
	r := &MovieRouters{t}

	handler.GET("/", r.getMovie)
	handler.GET("/:id", r.getMoviebyID)
	handler.POST("/", r.addMovie)
	handler.PATCH("/:id", r.updateMovie)
	handler.DELETE("/:id", r.deleteeMovie)
}

func (r *MovieRouters) getMovie(c *gin.Context) {
	res, err := r.m.GetMovie(c)
	log.Println("respon : ", res)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": res})
		return
	}

}
func (r *MovieRouters) getMoviebyID(c *gin.Context) {
	id := c.Param("id")
	res, err := r.m.GetMoviebyID(c, id)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
	return

}
func (r *MovieRouters) addMovie(c *gin.Context) {
	var req entity.Movies
	c.BindJSON(&req)
	err := r.m.AddMovie(c, req)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, "")
		return
	}
}
func (r *MovieRouters) updateMovie(c *gin.Context) {
	var req entity.Movies
	c.BindJSON(&req)
	err := r.m.UpdateMovie(c, req)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, "")
		return
	}
}
func (r *MovieRouters) deleteeMovie(c *gin.Context) {

	id := c.Param("id")
	err := r.m.DeleteeMovie(c, id)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, "")
		return
	}
}
