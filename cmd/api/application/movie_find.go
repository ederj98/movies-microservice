package application

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

const (
	errorServiceMovieFind = "error getting information from service Movie"
)

// FindAllApplication is the initial flow entry to granted the get to the movie
type MovieFindApplication interface {
	// Handler is the input for get the movie
	Handler(id int64) (movie model.Movie, err error)
}

type MovieFind struct {
	MovieFindService service.MovieFindServicePort
}

func (movieFind *MovieFind) Handler(id int64) (movie model.Movie, err error) {

	movie, err = movieFind.MovieFindService.Find(id)

	if err != nil {
		logger.Error(errorServiceMovieFind, err)
		return movie, err
	}
	return movie, err
}
