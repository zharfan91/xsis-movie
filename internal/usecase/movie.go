package usecase

import (
	"errors"
	"log"
	"strconv"
	"xsis-movie/internal/entity"

	"github.com/gin-gonic/gin"
)

type MovieUseCase struct {
	repo MovieRepo
}

func New(r MovieRepo) *MovieUseCase {
	return &MovieUseCase{
		repo: r,
	}
}

// addMovie implements Movies.
func (m *MovieUseCase) AddMovie(c *gin.Context, req entity.Movies) *entity.RequestError {
	err := m.repo.AddMovie(req)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return err
	}
	return nil
}

// deleteeMovie implements Movies.
func (m *MovieUseCase) DeleteeMovie(c *gin.Context, req string) *entity.RequestError {
	val, err := strconv.Atoi(req)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return err
	}
	err = m.repo.DeleteeMovie(val)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return err
	}
	return nil
}

// getMovie implements Movies.
func (m *MovieUseCase) GetMovie(c *gin.Context) ([]entity.Movies, *entity.RequestError) {
	res, err := m.repo.GetMovie()
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return nil, err
	}
	return res, nil
}

// getMoviebyID implements Movies.
func (m *MovieUseCase) GetMoviebyID(c *gin.Context, req string) (*entity.Movies, *entity.RequestError) {
	val, err := strconv.Atoi(req)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return nil, err
	}
	res, err := m.repo.GetMoviebyID(val)
	log.Println("GetMoviebyID respon : ", res)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return nil, err
	}
	return res, nil
}

// updateMovie implements Movies.
func (m *MovieUseCase) UpdateMovie(c *gin.Context, req entity.Movies) *entity.RequestError {
	err := m.repo.UpdateMovie(&req)
	if err != nil {
		err := &entity.RequestError{
			StatusCode: 503,
			Err:        errors.New("The server cannot handle the request"),
		}
		return err
	}
	return nil
}
