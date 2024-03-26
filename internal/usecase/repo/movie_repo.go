package repo

import (
	"log"
	"xsis-movie/internal/entity"

	"gorm.io/gorm"
)

type MovieRepo struct {
	pg *gorm.DB
}

func New(pg *gorm.DB) *MovieRepo {
	return &MovieRepo{
		pg: pg,
	}
}

// addMovie implements Movies.
func (m *MovieRepo) AddMovie(req entity.Movies) error {
	err := m.pg.Create(&req)
	if err != nil {
		return err.Error
	}
	log.Print("Error: ", err)
	return nil
}

// deleteeMovie implements Movies.
func (m *MovieRepo) DeleteeMovie(id int) error {
	err := m.pg.Delete(&entity.Movies{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// getMovie implements Movies.
func (m *MovieRepo) GetMovie() ([]entity.Movies, error) {
	var movies []entity.Movies
	rows := m.pg.Find(&movies)

	if rows.Error != nil {
		return nil, rows.Error
	}
	return movies, nil
}

// getMoviebyID implements Movies.
func (m *MovieRepo) GetMoviebyID(id int) (*entity.Movies, error) {
	log.Println("start : 1")
	var movies *entity.Movies
	m.pg.Where("id = ?", id).Find(&movies)
	return movies, nil
}

// updateMovie implements Movies.
func (m *MovieRepo) UpdateMovie(movie *entity.Movies) error {
	log.Println("movie : ", movie)
	err := m.pg.Updates(movie).Error
	if err != nil {
		return err
	}
	return nil
}
