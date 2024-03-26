package usecase

import (
	"xsis-movie/internal/entity"

	"github.com/gin-gonic/gin"
)

type (
	Movies interface {
		GetMovie(*gin.Context) ([]entity.Movies, *entity.RequestError)
		GetMoviebyID(*gin.Context, string) (*entity.Movies, *entity.RequestError)
		AddMovie(*gin.Context, entity.Movies) *entity.RequestError
		UpdateMovie(*gin.Context, entity.Movies) *entity.RequestError
		DeleteeMovie(*gin.Context, string) *entity.RequestError
	}
	MovieRepo interface {
		GetMovie() ([]entity.Movies, error)
		GetMoviebyID(int) (*entity.Movies, error)
		AddMovie(entity.Movies) error
		UpdateMovie(*entity.Movies) error
		DeleteeMovie(id int) error
	}
)
