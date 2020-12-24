package application

import (
	"fmt"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

// MovieDeleteApplication is the initial flow entry to granted the delete the movie
type MovieDeleteApplication interface {
	// Handler is the input for update the movie
	Handler(id int64) (err error)
}

type MovieDelete struct {
	MovieDeleteService service.MovieDeleteServicePort
}

func (movieDelete *MovieDelete) Handler(id int64) (err error) {

	err = movieDelete.MovieDeleteService.Delete(id)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, id), err)
		return err
	}
	return err
}
