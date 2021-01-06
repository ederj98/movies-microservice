package application

import (
	"fmt"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

// MovieUpdateApplication is the initial flow entry to granted the update  the movie
type MovieUpdateApplication interface {
	// Handler is the input for update the movie
	Handler(id int64, movie model.Movie) (err error)
}

type MovieUpdate struct {
	MovieUpdateService service.MovieUpdateServicePort
}

func (movieUpdate *MovieUpdate) Handler(id int64, movie model.Movie) (err error) {

	err = movieUpdate.MovieUpdateService.Update(id, movie)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, movie.Name), err)
		return err
	}
	return err
}
