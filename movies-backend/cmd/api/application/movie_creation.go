package application

import (
	"fmt"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"

	"github.com/ederj98/movies-microservice/pkg/logger"
)

const (
	errorServiceMovie = "error getting information from service Movie Creation for Movie %v"
)

// MovieCreationApplication is the initial flow entry to granted the create  the movie
type MovieCreationApplication interface {
	// Handler is the input for create the movie
	Handler(movie model.Movie) (err error)
}

type MovieCreation struct {
	MovieCreationService service.MovieCreationServicePort
}

func (movieCreation *MovieCreation) Handler(movie model.Movie) (err error) {

	err = movieCreation.MovieCreationService.Create(movie)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, movie.Name), err)
		return err
	}
	return err
}
