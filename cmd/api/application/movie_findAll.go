package application

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

const (
	errorServiceMovieList = "error getting information from service Movie List"
)

// FindAllApplication is the initial flow entry to granted the get to the movies
type MovieFindAllApplication interface {
	// Handler is the input for get the movies
	Handler() (movieLots []model.Movie, err error)
}

type MovieFindAll struct {
	MovieRepository port.MovieRepository
}

func (movieFindAll *MovieFindAll) Handler() (movieLots []model.Movie, err error) {

	movieLots, err = movieFindAll.MovieRepository.FindAll()

	if err != nil {
		logger.Error(errorServiceMovieList, err)
		return movieLots, err
	}
	return movieLots, err
}
